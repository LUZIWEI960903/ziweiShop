package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsCateLogic struct {
}

func (GoodsCateLogic) DoAdd(p *models.AddGoodsCateParams, cateImg string) (err error) {
	return mysql.AddGoodsCate(p, cateImg)
}
