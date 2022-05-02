package logic

import (
	mysql "ziweiShop/dao/mysql/admin"
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

func (GoodsColorLogic) ShowEditPageLogic(goodsColorId int) (goodsColorInfo *models.GoodsColorList, err error) {
	// 获取 商品颜色详情
	oGoodesColor, err := mysql.GetGoodsColorById(goodsColorId)
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	goodsColorInfo = &models.GoodsColorList{
		Id:         oGoodesColor.Id,
		Status:     oGoodesColor.Status,
		ColorName:  oGoodesColor.ColorName,
		ColorValue: oGoodesColor.ColorValue,
	}
	return goodsColorInfo, nil
}

func (GoodsColorLogic) DoEditGoodsColorLogic(p *models.EditGoodsColorParams) (err error) {
	return mysql.EditGoodsColor(p)
}

func (GoodsColorLogic) DeleteGoodsColorLogic(goodsColorId int) (err error) {
	return mysql.DeleteGoodsColor(goodsColorId)
}
