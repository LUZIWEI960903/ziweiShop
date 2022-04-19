package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsColorLogic struct {
}

func (GoodsColorLogic) DoAddGoodsColorLogic(p *models.AddGoodsColorParams) (err error) {
	return mysql.AddGoodsColor(p)
}
