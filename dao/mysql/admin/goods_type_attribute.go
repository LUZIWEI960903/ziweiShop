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

// EditGoodsTypeAttribute 根据 goodsTypeAttributeId 修改 商品类型属性  --- goods_type_attribute 表
func EditGoodsTypeAttribute(p *models.EditGoodsTypeAttributeParams) (err error) {
	oGoodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err = db.Where("id=? AND is_deleted=0", p.Id).First(&oGoodsTypeAttributeList).Error
	if len(oGoodsTypeAttributeList) < 1 || err != nil {
		return ErrGetGoodsTypeAttribute
	}
	oGoodsTypeAttributeList[0].CateId = p.CateId
	oGoodsTypeAttributeList[0].Status = p.Status
	oGoodsTypeAttributeList[0].Sort = p.Sort
	oGoodsTypeAttributeList[0].AttrType = p.AttrType
	oGoodsTypeAttributeList[0].Title = p.Title
	oGoodsTypeAttributeList[0].AttrValue = p.AttrValue
	return db.Save(&oGoodsTypeAttributeList).Error
}

// DeleteGoodsTypeAttribute 根据 goodsTypeAttributeId 逻辑删除 商品类型属性  --- goods_type_attribute 表
func DeleteGoodsTypeAttribute(goodsTypeAttributeId int) (err error) {
	oGoodsTypeAttributeList := []models.GoodsTypeAttribute{}
	err = db.Where("id=? AND is_deleted=0", goodsTypeAttributeId).First(&oGoodsTypeAttributeList).Error
	if len(oGoodsTypeAttributeList) < 1 || err != nil {
		return ErrGetGoodsTypeAttribute
	}

	oGoodsTypeAttributeList[0].IsDeleted = 1
	return db.Save(&oGoodsTypeAttributeList).Error
}

// GetGoodsTypeAttributeByCateId 根据 cateId 获取所有的 商品类型属性  --- goods_type_attribute 表
func GetGoodsTypeAttributeByCateId(cateId int) (goodsTypeAttributeList []models.GoodsTypeAttribute, err error) {
	goodsTypeAttributeList = []models.GoodsTypeAttribute{}
	err = db.Where("cate_id=? AND is_deleted=0", cateId).Find(&goodsTypeAttributeList).Error
	if len(goodsTypeAttributeList) < 1 || err != nil {
		return nil, ErrGetGoodsTypeAttribute
	}
	return goodsTypeAttributeList, err
}
