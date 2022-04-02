package admin

import (
	"strings"
	"ziweiShop/models"
)

func verifyAddRoleParams(p *models.AddRoleParams) (err error) {
	title := strings.Trim(p.Title, " ")
	if title == "" {
		return ErrEmptyTitle
	}
	return nil
}

func verifyEditRoleParams(p *models.EditRoleParams) (err error) {
	title := strings.Trim(p.Title, " ")
	if title == "" {
		return ErrEmptyTitle
	}
	return nil
}
