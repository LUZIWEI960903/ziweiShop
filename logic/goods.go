package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsLogic struct {
}

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
