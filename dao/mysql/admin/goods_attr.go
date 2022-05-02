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

// DeleteGoodsAttr  根据 goodsId 逻辑删除 对应的包装规格  --- goods_attr 表
func DeleteGoodsAttr(goodsId int) (err error) {
	goodsAttrList := []models.GoodsAttr{}
	err = db.Where("goods_id=? AND is_deleted=0", goodsId).Find(&goodsAttrList).Error
	if err != nil {
		return err
	}

	for i, l := 0, len(goodsAttrList); i < l; i++ { // 这里不能使用range遍历，range遍历会拷贝对象，不会修改原对象
		goodsAttrList[i].IsDeleted = 1
	}
	return db.Save(&goodsAttrList).Error
}
