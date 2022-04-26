package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type NavLogic struct {
}

func (NavLogic) AddNavLogic(p *models.AddNavParams) error {
	return mysql.AddNav(p)
}
