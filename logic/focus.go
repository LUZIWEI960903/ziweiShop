package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type FocusLogic struct {
}

func (FocusLogic) DoAdd(p *models.AddFocusParams, focusImgSrc string) (err error) {
	return mysql.AddFocus(p, focusImgSrc)
}
