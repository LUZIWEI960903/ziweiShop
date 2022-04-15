package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddGoodsTypeAttribute 添加 商品类型属性   --- goods_type_attribute 表
func AddGoodsTypeAttribute(p *models.AddGoodsTypeAttributeParams) (err error) {
	goodsTypeAttribute := models.GoodsTypeAttribute{
		CateId:    p.CateId,
		Status:    p.Status,
		Sort:      p.Sort,
		AddTime:   int(tools.GetUnix()),
		AttrType:  p.AttrType,
		Title:     p.Title,
		AttrValue: p.AttrValue,
	}
	return db.Create(&goodsTypeAttribute).Error
}

// GetGoodsTypeAttributeList 根据 cateId 查询其下的 所有商品类型属性   --- goods_type_attribute 表
func GetGoodsTypeAttributeList(cateId int) (oGoodsTypeAttributeList []models.GoodsTypeAttribute, err error) {
	oGoodsTypeAttributeList = []models.GoodsTypeAttribute{}
	err = db.Where("cate_id=? AND is_deleted=0", cateId).Find(&oGoodsTypeAttributeList).Error
	return oGoodsTypeAttributeList, err
}
