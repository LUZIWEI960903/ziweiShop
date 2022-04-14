package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddGoodsType 增加商品类型  --- goods_type 表
func AddGoodsType(p *models.AddGoodsTypeParams) (err error) {
	goodsType := models.GoodsType{
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		AddTime:     int(tools.GetUnix()),
	}
	return db.Create(&goodsType).Error
}

// GetGoodsTypeList 查询所有商品类型  --- goods_type 表
func GetGoodsTypeList() (oGoodsTypeList []models.GoodsType, err error) {
	oGoodsTypeList = []models.GoodsType{}
	err = db.Where("is_deleted=0").Find(&oGoodsTypeList).Error
	if len(oGoodsTypeList) < 1 {
		return nil, ErrGetGoodsType
	}
	return oGoodsTypeList, err
}

// GetGoodsTypeById 根据 goodsTypeId 查询 商品类型信息 --- goods_type 表
func GetGoodsTypeById(goodsTypeId int) (oGoodsTypeInfo *models.GoodsType, err error) {
	oGoodsTypeInfoList := []models.GoodsType{}
	err = db.Where("id=? AND is_deleted=0", goodsTypeId).First(&oGoodsTypeInfoList).Error
	if len(oGoodsTypeInfoList) < 1 {
		return nil, ErrGetGoodsType
	}
	return &oGoodsTypeInfoList[0], err
}

// EditGoodsType 修改 商品类型信息  --- goods_type 表
func EditGoodsType(p *models.EditGoodsTypeParams) (err error) {
	goodTypeList := []models.GoodsType{}
	db.Where("id=? AND is_deleted=0", p.Id).First(&goodTypeList)
	if len(goodTypeList) < 1 {
		return ErrGetGoodsType
	}
	goodTypeList[0].Title = p.Title
	goodTypeList[0].Description = p.Description
	goodTypeList[0].Status = p.Status

	return db.Save(&goodTypeList).Error
}

// DeleteGoodsType 逻辑删除删除 商品类型  --- goods_type 表
func DeleteGoodsType(goodsTypeId int) (err error) {
	goodsTypeList := []models.GoodsType{}
	db.Where("id=? AND is_deleted=0", goodsTypeId).First(&goodsTypeList)
	if len(goodsTypeList) < 1 {
		return ErrGetGoodsType
	}
	goodsTypeList[0].IsDeleted = 1
	return db.Save(&goodsTypeList).Error
}
