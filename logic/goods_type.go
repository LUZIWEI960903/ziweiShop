package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsTypeLogic struct {
}

func (GoodsTypeLogic) DoAdd(p *models.AddGoodsTypeParams) (err error) {
	return mysql.AddGoodsType(p)
}

func (GoodsTypeLogic) GetGoodsTypeList() (goodsTypeList []models.GetGoodsType, err error) {
	// 获取商品种类列表
	oGoodsTypeList, err := mysql.GetGoodsTypeList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	for _, oGoodsType := range oGoodsTypeList {
		goodType := models.GetGoodsType{
			Id:          oGoodsType.Id,
			Status:      oGoodsType.Status,
			Title:       oGoodsType.Title,
			Description: oGoodsType.Description,
		}
		goodsTypeList = append(goodsTypeList, goodType)
	}

	return goodsTypeList, nil
}
