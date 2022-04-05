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

func verifyAddAccessParams(p *models.AddAccessParams) (err error) {
	moduleName := strings.Trim(p.ModuleName, " ")
	if moduleName == "" {
		return ErrEmptyModuleName
	}
	return nil
}

func verifyEditAccessParams(p *models.EditAccessParams) (err error) {
	moduleName := strings.Trim(p.ModuleName, " ")
	if moduleName == "" {
		return ErrEmptyModuleName
	}
	return nil
}
