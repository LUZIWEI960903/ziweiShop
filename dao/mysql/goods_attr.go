package mysql

import "ziweiShop/models"

// AddGoodsAttr 创建 goodsAttr   --- goods_attr 表
func AddGoodsAttr(goodsAttrObj *models.GoodsAttr) error {
	return db.Create(&goodsAttrObj).Error
}
