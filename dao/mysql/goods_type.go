package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddGoodsType 增加商品类型
func AddGoodsType(p *models.AddGoodsTypeParams) (err error) {
	goodsType := models.GoodsType{
		Title:       p.Title,
		Description: p.Description,
		Status:      p.Status,
		AddTime:     int(tools.GetUnix()),
	}
	return db.Create(&goodsType).Error
}
