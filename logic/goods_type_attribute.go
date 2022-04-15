package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

type GoodsTypeAttributeLogic struct {
}

func (GoodsTypeAttributeLogic) ShowAddPageLogic(cateId int) (data *models.GoodsTypeAttributeAddPageData, err error) {
	// 查询 cateId 是否存在
	_, err = mysql.GetGoodsCateById(cateId)
	if err != nil {
		return nil, err
	}

	// 查询所有 商品类型
	oGoodsTypeList, err := mysql.GetGoodsTypeList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	GoodsTypeItems := make([]models.GoodsTypeInfo, 0)
	GoodsTypeItems1 := make([]models.GoodsTypeInfo, 0)
	GoodsTypeItems2 := make([]models.GoodsTypeInfo, 0)
	for _, oGoodsType := range oGoodsTypeList {
		GoodsType := models.GoodsTypeInfo{
			Id:          oGoodsType.Id,
			Status:      oGoodsType.Status,
			Title:       oGoodsType.Title,
			Description: oGoodsType.Description,
		}
		if oGoodsType.Id == cateId {
			GoodsTypeItems1 = append(GoodsTypeItems1, GoodsType)
		} else {
			GoodsTypeItems2 = append(GoodsTypeItems2, GoodsType)
		}
	}
	GoodsTypeItems = append(GoodsTypeItems1, GoodsTypeItems2...)

	return &models.GoodsTypeAttributeAddPageData{
		CateId:         cateId,
		GoodsTypeItems: GoodsTypeItems,
	}, nil
}

func (GoodsTypeAttributeLogic) DoAddLogic(p *models.AddGoodsTypeAttributeParams) (err error) {
	// 查看 cateId 是否存在
	_, err = mysql.GetGoodsTypeById(p.CateId)
	if err != nil {
		return err
	}
	return mysql.AddGoodsTypeAttribute(p)
}

func (GoodsTypeAttributeLogic) ShowIndexPageLogic(cateId int) (data *models.GoodsTypeAttributeIndexPageData, err error) {
	// 查看 cateId 是否存在
	oGoodTypeInfo, err := mysql.GetGoodsTypeById(cateId)
	if err != nil {
		return nil, err
	}

	// 根据 cateId 查询其所有 商品类型属性
	oGoodsTypeAttributeList, err := mysql.GetGoodsTypeAttributeList(cateId)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	GoodsTypeAttributeList := make([]models.GoodsTypeAttributeInfo, 0)
	for _, oGoodsTypeAttribute := range oGoodsTypeAttributeList {
		GoodsTypeAttribute := models.GoodsTypeAttributeInfo{
			Id:        oGoodsTypeAttribute.Id,
			Status:    oGoodsTypeAttribute.Status,
			Sort:      oGoodsTypeAttribute.Sort,
			AttrType:  oGoodsTypeAttribute.AttrType,
			Title:     oGoodsTypeAttribute.Title,
			AttrValue: oGoodsTypeAttribute.AttrValue,
			AddTime:   tools.UnixToDate(oGoodsTypeAttribute.AddTime),
		}
		GoodsTypeAttributeList = append(GoodsTypeAttributeList, GoodsTypeAttribute)
	}

	return &models.GoodsTypeAttributeIndexPageData{
		CateId:                      cateId,
		Title:                       oGoodTypeInfo.Title,
		GoodsTypeAttributeInfoItems: GoodsTypeAttributeList,
	}, nil
}

func (GoodsTypeAttributeLogic) ShowEditPageLogic(goodsTypeAttributeId int) (data *models.GoodsTypeAttributeEditPageData, err error) {
	// 查询 goodsTypeAttribute 是否存在
	oGoodsTypeAttribute, err := mysql.GetGoodsTypeAttributeById(goodsTypeAttributeId)
	if err != nil {
		return nil, err
	}

	// 查询 所有商品类型
	oGoodsTypeList, err := mysql.GetGoodsTypeList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	goodsTypeList := make([]models.GoodsTypeItems, 0)
	goodsTypeList1 := make([]models.GoodsTypeItems, 0)
	goodsTypeList2 := make([]models.GoodsTypeItems, 0)

	for _, oGoodsType := range oGoodsTypeList {
		goodsType := models.GoodsTypeItems{
			Id:    oGoodsType.Id,
			Title: oGoodsType.Title,
		}
		if oGoodsType.Id == oGoodsTypeAttribute.CateId {
			goodsTypeList1 = append(goodsTypeList1, goodsType)
		} else {
			goodsTypeList2 = append(goodsTypeList2, goodsType)
		}
	}
	goodsTypeList = append(goodsTypeList1, goodsTypeList2...)

	return &models.GoodsTypeAttributeEditPageData{
		Id:             oGoodsTypeAttribute.Id,
		Status:         oGoodsTypeAttribute.Status,
		Sort:           oGoodsTypeAttribute.Sort,
		AttrType:       oGoodsTypeAttribute.AttrType,
		Title:          oGoodsTypeAttribute.Title,
		AttrValue:      oGoodsTypeAttribute.AttrValue,
		GoodsTypeItems: goodsTypeList,
	}, nil
}
