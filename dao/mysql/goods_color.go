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

// GetGoodsColorList 获取所有 商品颜色  --- goods_color 表
func GetGoodsColorList() (oGoodsColorList []models.GoodsColor, err error) {
	oGoodsColorList = []models.GoodsColor{}
	err = db.Where("is_deleted=0").Find(&oGoodsColorList).Error
	return oGoodsColorList, err
}
