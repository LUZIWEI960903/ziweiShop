package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsColorLogic struct {
}

func (GoodsColorLogic) DoAddGoodsColorLogic(p *models.AddGoodsColorParams) (err error) {
	return mysql.AddGoodsColor(p)
}

func (GoodsColorLogic) ShowIndexPageLogic() (data []models.GoodsColorList, err error) {
	// 获取 所有商品颜色
	oGoodsColorList, err := mysql.GetGoodsColorList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	for _, oGoodsColor := range oGoodsColorList {
		newGoodsColor := models.GoodsColorList{
			Id:         oGoodsColor.Id,
			Status:     oGoodsColor.Status,
			ColorName:  oGoodsColor.ColorName,
			ColorValue: oGoodsColor.ColorValue,
		}
		data = append(data, newGoodsColor)
	}
	return data, nil
}
