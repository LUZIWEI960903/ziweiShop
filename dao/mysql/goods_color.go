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

// GetGoodsColorById 根据 goodsColorId 获取 商品颜色id
func GetGoodsColorById(goodsColorId int) (oGoodsColor *models.GoodsColor, err error) {
	oGoodsColorList := []models.GoodsColor{}
	err = db.Where("id=? AND is_deleted=0", goodsColorId).First(&oGoodsColorList).Error
	if len(oGoodsColorList) < 1 || err != nil {
		return nil, ErrGetGoodsColor
	}
	return &oGoodsColorList[0], nil
}
