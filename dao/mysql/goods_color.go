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

// GetGoodsColorById 根据 goodsColorId 获取 商品颜色id  --- goods_color 表
func GetGoodsColorById(goodsColorId int) (oGoodsColor *models.GoodsColor, err error) {
	oGoodsColorList := []models.GoodsColor{}
	err = db.Where("id=? AND is_deleted=0", goodsColorId).First(&oGoodsColorList).Error
	if len(oGoodsColorList) < 1 || err != nil {
		return nil, ErrGetGoodsColor
	}
	return &oGoodsColorList[0], nil
}

// GetGoodsColorById2 根据 goodsColorId 获取 商品颜色id  --- goods_color 表
func GetGoodsColorById2(goodsColorId int) (goodsColor *models.GoodsColor, err error) {
	return goodsColor, db.Where("id=? AND is_deleted=0", goodsColorId).Select("id,color_name").First(&goodsColor).Error
}

// EditGoodsColor 根据 goodsColorId 修改 商品颜色  --- goods_color 表
func EditGoodsColor(p *models.EditGoodsColorParams) (err error) {
	goodsColorList := []models.GoodsColor{}
	err = db.Where("id=? AND is_deleted=0", p.Id).First(&goodsColorList).Error
	if len(goodsColorList) < 1 || err != nil {
		return ErrGetGoodsColor
	}
	goodsColorList[0].Status = p.Status
	goodsColorList[0].ColorName = p.ColorName
	goodsColorList[0].ColorValue = p.ColorValue
	return db.Save(&goodsColorList).Error
}

// DeleteGoodsColor 根据 goodsColorId 逻辑删除 商品颜色  --- goods_color 表
func DeleteGoodsColor(goodsColorId int) (err error) {
	goodsColorList := []models.GoodsColor{}
	err = db.Where("id=? AND is_deleted=0", goodsColorId).First(&goodsColorList).Error
	if len(goodsColorList) < 1 || err != nil {
		return ErrGetGoodsColor
	}
	goodsColorList[0].IsDeleted = 1
	return db.Save(&goodsColorList).Error
}
