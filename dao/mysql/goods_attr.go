package mysql

import "ziweiShop/models"

// AddGoodsAttr 创建 goodsAttr   --- goods_attr 表
func AddGoodsAttr(goodsAttrObj *models.GoodsAttr) error {
	return db.Create(&goodsAttrObj).Error
}

// GetGoodsAttrListByGoodsId  根据 goodsId  查询 其下的所有 GoodsAttr  --- goods_attr 表
func GetGoodsAttrListByGoodsId(goodsId int) (oGoodsAttrList []models.GoodsAttr, err error) {
	oGoodsAttrList = []models.GoodsAttr{}
	err = db.Where("goods_id=? AND is_deleted=0", goodsId).Find(&oGoodsAttrList).Error
	return oGoodsAttrList, err
}
