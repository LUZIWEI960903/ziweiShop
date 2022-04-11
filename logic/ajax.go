package logic

import "ziweiShop/dao/mysql"

type AjaxLogic struct {
}

func (AjaxLogic) ChangeStatus(id, table, field string) (err error) {
	return mysql.AjaxChangeStatus(id, table, field)
}
