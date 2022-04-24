package logic

import (
	"strconv"
	"strings"
	"sync"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

type GoodsLogic struct {
}

var wg sync.WaitGroup

func (GoodsLogic) ShowAddPageLogic() (data *models.GoodsAddPageData, err error) {
	// 获取商品分类
	oTopGoodsCateWithGoodsCateList, err := mysql.GetTopGoodsCateWithGoodsCateList()
	if err != nil {
		return nil, err
	}
	// 获取商品颜色信息
	oGoodsColorList, err := mysql.GetGoodsColorList()
	if err != nil {
		return nil, err
	}
	// 获取商品规格包装
	oGoodsTypeList, err := mysql.GetGoodsTypeList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	GoodsCateItems := make([]models.GoodsCateWithGoodsCate1, 0)
	for _, TopGoodsCate := range oTopGoodsCateWithGoodsCateList {
		newTopGoodsCate := models.GoodsCateWithGoodsCate1{
			Id:    TopGoodsCate.Id,
			Pid:   TopGoodsCate.Pid,
			Title: TopGoodsCate.Title,
			//ChildGoodsCateItems: nil,
		}
		ChildGoodsCateItems := make([]models.TopGoodsCate2, 0)
		for _, GoodsCate := range TopGoodsCate.GoodsCateItems {
			newGoodsCate := models.TopGoodsCate2{
				Id:    GoodsCate.Id,
				Pid:   GoodsCate.Pid,
				Title: GoodsCate.Title,
			}
			ChildGoodsCateItems = append(ChildGoodsCateItems, newGoodsCate)
		}
		newTopGoodsCate.ChildGoodsCateItems = ChildGoodsCateItems
		GoodsCateItems = append(GoodsCateItems, newTopGoodsCate)
	}

	GoodsColorItems := make([]models.GoodsColorList1, 0)
	for _, oGoodsColor := range oGoodsColorList {
		newGoodsColor := models.GoodsColorList1{
			Id:        oGoodsColor.Id,
			ColorName: oGoodsColor.ColorName,
		}
		GoodsColorItems = append(GoodsColorItems, newGoodsColor)
	}

	GoodsTypeItems := make([]models.GoodsTypeList1, 0)
	for _, oGoodsType := range oGoodsTypeList {
		newGoodsType := models.GoodsTypeList1{
			Id:    oGoodsType.Id,
			Title: oGoodsType.Title,
		}
		GoodsTypeItems = append(GoodsTypeItems, newGoodsType)
	}

	return &models.GoodsAddPageData{
		GoodsCateItems:  GoodsCateItems,
		GoodsColorItems: GoodsColorItems,
		GoodsTypeItems:  GoodsTypeItems,
	}, nil
}

func (GoodsLogic) AjaxGetGoodsTypeAttributeLogic(cateId int) (data []models.AjaxGoodsTypeAttribute, err error) {
	// 根据 查询所有 满足该 cateId的 商品类型属性
	oGoodsTypeAttributeList, err := mysql.GetGoodsTypeAttributeByCateId(cateId)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	for _, oGoodsTypeAttribute := range oGoodsTypeAttributeList {
		newGoodsTypeAttribute := models.AjaxGoodsTypeAttribute{
			Id:        oGoodsTypeAttribute.Id,
			CateId:    oGoodsTypeAttribute.CateId,
			Title:     oGoodsTypeAttribute.Title,
			AttrType:  oGoodsTypeAttribute.AttrType,
			AttrValue: oGoodsTypeAttribute.AttrValue,
		}
		data = append(data, newGoodsTypeAttribute)
	}
	return data, nil
}

func (GoodsLogic) AddGoodsLogic(p *models.AddGoodsParams, goodsImageList, attrIdList, attrValueList []string) (err error) {
	// 创建商品
	goods, err := mysql.AddGoods(p)
	if err != nil {
		return err
	}

	// 增加图库信息
	wg.Add(1)
	go func() error {
		for _, goodsImageUrl := range goodsImageList {
			goodsImageObj := models.GoodsImage{
				GoodsId: goods.Id,
				ColorId: 0,
				Sort:    10,
				AddTime: int(tools.GetUnix()),
				Status:  1,
				ImgUrl:  goodsImageUrl,
			}
			if err1 := mysql.AddGoodsImage(&goodsImageObj); err1 != nil {
				return err1
			}
		}
		wg.Done()
		return nil
	}()

	// 增加规格包装
	wg.Add(1)
	go func() error {
		for i, l := 0, len(attrIdList); i < l; i++ {
			// 获取商品类型属性 信息
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			oGoodsTypeAttribute, err1 := mysql.GetGoodsTypeAttributeById(goodsTypeAttributeId)
			if err1 != nil {
				return err1
			}
			// 创建 goodsAttr
			goodsAttrObj := models.GoodsAttr{
				GoodsId:         goods.Id,
				AttributeCateId: oGoodsTypeAttribute.CateId,
				AttributeId:     oGoodsTypeAttribute.Id,
				AttributeType:   oGoodsTypeAttribute.AttrType,
				AddTime:         int(tools.GetUnix()),
				Status:          1,
				AttributeTitle:  oGoodsTypeAttribute.Title,
				AttributeValue:  attrValueList[i],
				Sort:            10,
			}
			if err2 := mysql.AddGoodsAttr(&goodsAttrObj); err2 != nil {
				return err2
			}
		}
		wg.Done()
		return nil
	}()

	wg.Wait()

	return nil
}

func (GoodsLogic) ShowIndexPageDataLogic() (data *models.GoodsIndexPageData, err error) {
	// 查询所有商品
	oGoodsList, err := mysql.GetGoodsList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	GoodsItems := make([]models.GoodsIndexPage, 0)
	for _, oGoods := range oGoodsList {
		newGoods := models.GoodsIndexPage{
			Id:          oGoods.Id,
			Title:       oGoods.Title,
			Price:       oGoods.Price,
			MarketPrice: oGoods.MarketPrice,
			IsHot:       oGoods.IsHot,
			IsBest:      oGoods.IsBest,
			IsNew:       oGoods.IsNew,
			ClickCount:  oGoods.ClickCount,
			GoodsNumber: oGoods.GoodsNumber,
			Sort:        oGoods.Sort,
			Status:      oGoods.Status,
		}
		GoodsItems = append(GoodsItems, newGoods)
	}

	return &models.GoodsIndexPageData{
		GoodsItems: GoodsItems,
	}, nil
}

func (GoodsLogic) ShowEditPageLogic(goodsId int) (data *models.GoodsEditPageData, err error) {
	// 查询 goods 信息
	oGoodsInfo, err1 := mysql.GetGoodsById(goodsId)
	if err1 != nil {
		return nil, err1
	}

	// 查询所有 该商品图库信息
	oGoodsImageList, err2 := mysql.GetGoodsImageListByGoodsId(goodsId)
	if err2 != nil {
		return nil, err2
	}

	// 查询所有 该商品的包装规格信息
	oGoodsAttrList, err3 := mysql.GetGoodsAttrListByGoodsId(goodsId)
	if err3 != nil {
		return nil, err3
	}

	// 查询所有 商品分类
	oTopGoodsCateWithGoodsCateList, err4 := mysql.GetTopGoodsCateWithGoodsCateList()
	if err4 != nil {
		return nil, err4
	}

	// 查询所有 商品颜色
	oGoodsColorList, err5 := mysql.GetGoodsColorList()
	if err5 != nil {
		return nil, err5
	}

	// 查询所有 商品类型
	oGoodsTypeList, err6 := mysql.GetGoodsTypeList()
	if err6 != nil {
		return nil, err6
	}

	// 构造返回数据
	Goods := models.GoodsEdit{
		CateId:        oGoodsInfo.CateId,
		GoodsNumber:   oGoodsInfo.GoodsNumber,
		IsHot:         oGoodsInfo.IsHot,
		IsBest:        oGoodsInfo.IsBest,
		IsNew:         oGoodsInfo.IsNew,
		GoodsTypeId:   oGoodsInfo.GoodsTypeId,
		Sort:          oGoodsInfo.Sort,
		Status:        oGoodsInfo.Status,
		Price:         oGoodsInfo.Price,
		MarketPrice:   oGoodsInfo.MarketPrice,
		Title:         oGoodsInfo.Title,
		SubTitle:      oGoodsInfo.SubTitle,
		GoodsSn:       oGoodsInfo.GoodsSn,
		RelationGoods: oGoodsInfo.RelationGoods,
		GoodsAttr:     oGoodsInfo.GoodsAttr,
		GoodsVersion:  oGoodsInfo.GoodsVersion,
		GoodsImg:      oGoodsInfo.GoodsImg,
		GoodsGift:     oGoodsInfo.GoodsGift,
		GoodsFitting:  oGoodsInfo.GoodsFitting,
		GoodsKeywords: oGoodsInfo.GoodsKeywords,
		GoodsDesc:     oGoodsInfo.GoodsDesc,
		GoodsContent:  oGoodsInfo.GoodsContent,
	}

	GoodsImageItems := make([]models.GoodsImageList, 0)
	for _, oGoodsImage := range oGoodsImageList {
		newGoodsImage := models.GoodsImageList{
			Id:      oGoodsImage.Id,
			GoodsId: oGoodsImage.GoodsId,
			ImgUrl:  oGoodsImage.ImgUrl,
		}
		GoodsImageItems = append(GoodsImageItems, newGoodsImage)
	}

	GoodsAttrItems := make([]models.GoodsAttrList, 0)
	for _, oGoodsAttr := range oGoodsAttrList {
		newGoodsAttr := models.GoodsAttrList{
			Id:              oGoodsAttr.Id,
			GoodsId:         oGoodsAttr.GoodsId,
			AttributeCateId: oGoodsAttr.AttributeCateId,
			AttributeId:     oGoodsAttr.AttributeId,
			AttributeType:   oGoodsAttr.AttributeType,
			AttributeTitle:  oGoodsAttr.AttributeTitle,
			AttributeValue:  oGoodsAttr.AttributeValue,
		}
		GoodsAttrItems = append(GoodsAttrItems, newGoodsAttr)
	}

	GoodsCateItems := make([]models.GoodsCateWithGoodsCate1, 0)
	for _, oTopGoodsCate := range oTopGoodsCateWithGoodsCateList {
		newTopGoodsCate := models.GoodsCateWithGoodsCate1{
			Id:    oTopGoodsCate.Id,
			Pid:   oTopGoodsCate.Pid,
			Title: oTopGoodsCate.Title,
			//ChildGoodsCateItems: nil,
		}
		ChildGoodsCateItems := make([]models.TopGoodsCate2, 0)
		for _, oGoodsCate := range oTopGoodsCate.GoodsCateItems {
			newGoodsCate := models.TopGoodsCate2{
				Id:    oGoodsCate.Id,
				Pid:   oGoodsCate.Pid,
				Title: oGoodsCate.Title,
			}
			ChildGoodsCateItems = append(ChildGoodsCateItems, newGoodsCate)
		}
		newTopGoodsCate.ChildGoodsCateItems = ChildGoodsCateItems
		GoodsCateItems = append(GoodsCateItems, newTopGoodsCate)
	}

	selectedColorIdList := strings.Split(oGoodsInfo.GoodsColor, ",")
	goodsColorIdMap := make(map[string]bool)
	for _, colorIdStr := range selectedColorIdList {
		goodsColorIdMap[colorIdStr] = true
	}
	GoodsColorItems := make([]models.GoodsColorList2, 0)
	for _, oGoodsColor := range oGoodsColorList {
		newGoodsColor := models.GoodsColorList2{
			Id:        oGoodsColor.Id,
			ColorName: oGoodsColor.ColorName,
		}
		if _, ok := goodsColorIdMap[strconv.Itoa(oGoodsColor.Id)]; ok {
			newGoodsColor.IsCheck = true
		}
		GoodsColorItems = append(GoodsColorItems, newGoodsColor)
	}

	GoodsTypeItems := make([]models.GoodsTypeList1, 0)
	for _, oGoodsType := range oGoodsTypeList {
		newGoodsType := models.GoodsTypeList1{
			Id:    oGoodsType.Id,
			Title: oGoodsType.Title,
		}
		GoodsTypeItems = append(GoodsTypeItems, newGoodsType)
	}

	return &models.GoodsEditPageData{
		Goods:           Goods,
		GoodsCateItems:  GoodsCateItems,
		GoodsColorItems: GoodsColorItems,
		GoodsTypeItems:  GoodsTypeItems,
		GoodsImageItems: GoodsImageItems,
		GoodsAttrItems:  GoodsAttrItems,
	}, nil
}

func (GoodsLogic) EditGoodsLogic(p *models.EditGoodsParams, goodsImageList, attrIdList, attrValueList []string) (err error) {
	// 根据 goodsId 修改商品 信息
	err1 := mysql.EditGoods(p)
	if err1 != nil {
		return err1
	}

	// 增加商品图库信息
	wg.Add(1)
	go func() error {
		for _, goodsImageUrl := range goodsImageList {
			goodsImageObj := models.GoodsImage{
				GoodsId: p.Id,
				ColorId: 0,
				Sort:    10,
				AddTime: int(tools.GetUnix()),
				Status:  1,
				ImgUrl:  goodsImageUrl,
			}
			if err2 := mysql.AddGoodsImage(&goodsImageObj); err2 != nil {
				return err2
			}
		}
		wg.Done()
		return nil
	}()

	// 增加规格包装
	wg.Add(1)
	go func() error {
		// 删除原有的规格包装
		if err := mysql.DeleteGoodsAttr(p.Id); err != nil {
			return err
		}

		for i, l := 0, len(attrIdList); i < l; i++ {
			// 获取商品类型属性 信息
			goodsTypeAttributeId, _ := strconv.Atoi(attrIdList[i])
			oGoodsTypeAttribute, err3 := mysql.GetGoodsTypeAttributeById(goodsTypeAttributeId)
			if err3 != nil {
				return err3
			}
			// 创建 goodsAttr
			goodsAttrObj := models.GoodsAttr{
				GoodsId:         p.Id,
				AttributeCateId: oGoodsTypeAttribute.CateId,
				AttributeId:     oGoodsTypeAttribute.Id,
				AttributeType:   oGoodsTypeAttribute.AttrType,
				AddTime:         int(tools.GetUnix()),
				Status:          1,
				AttributeTitle:  oGoodsTypeAttribute.Title,
				AttributeValue:  attrValueList[i],
				Sort:            10,
			}
			if err4 := mysql.AddGoodsAttr(&goodsAttrObj); err4 != nil {
				return err4
			}
		}
		wg.Done()
		return nil
	}()

	wg.Wait()

	return nil
}

func (GoodsLogic) AjaxChangeGoodsImageLogic(goodsImageId, colorId int) (err error) {
	// 根据 goodsImageId 修改colorId
	return mysql.EditColorByGoodsImageId(goodsImageId, colorId)
}

func (GoodsLogic) AjaxRemoveGoodsImageLogic(goodsImageId int) (err error) {
	return mysql.AjaxRemoveGoodsImageByGoodsImageId(goodsImageId)
}
