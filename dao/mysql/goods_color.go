package mysql

import "ziweiShop/models"

// AddGoodsColor 增加商品颜色  --- goods_color 表
func AddGoodsColor(p *models.AddGoodsColorParams) (err error) {
	goodsColor := models.GoodsColor{
		Status:     1,
		ColorName:  p.ColorName,
		ColorValue: p.ColorValue,
	}
	return db.Create(&goodsColor).Error
}
