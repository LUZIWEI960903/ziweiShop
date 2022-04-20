package mysql

import "ziweiShop/models"

// AddGoodsImage 创建 goodsImage   --- goods_image 表
func AddGoodsImage(goodsImageObj *models.GoodsImage) error {
	return db.Create(goodsImageObj).Error
}
