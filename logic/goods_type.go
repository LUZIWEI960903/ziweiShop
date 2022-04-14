package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsTypeLogic struct {
}

func (GoodsTypeLogic) DoAdd(p *models.AddGoodsTypeParams) (err error) {
	return mysql.AddGoodsType(p)
}
