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
