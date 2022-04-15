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

// GetGoodsTypeAttributeById 根据 goodsTypeAttributeId 查询 商品类型属性是否存在  --- goods_type_attribute 表
func GetGoodsTypeAttributeById(goodsTypeAttributeId int) (oGoodsTypeAttribute *models.GoodsTypeAttribute, err error) {
	oGoodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err = db.Where("id=? AND is_deleted=0", goodsTypeAttributeId).First(&oGoodsTypeAttributeList).Error
	if len(oGoodsTypeAttributeList) < 1 || err != nil {
		return nil, ErrGetGoodsTypeAttribute
	}
	return &oGoodsTypeAttributeList[0], nil
}
