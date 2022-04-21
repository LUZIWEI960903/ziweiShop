package logic

import (
	"strconv"
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
