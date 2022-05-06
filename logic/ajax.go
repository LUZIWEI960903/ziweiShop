package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/dao/redis"
)

type AjaxLogic struct {
}

func (AjaxLogic) ChangeStatus(id, table, field string) (err error) {
	return mysql.AjaxChangeStatus(id, table, field)
}

func (AjaxLogic) ChangeSort(id, table, field, num string) (err error) {
	return mysql.AjaxChangeSort(id, table, field, num)
}

func (AjaxLogic) FlushAll() {
	redis.FlushAll()
}
