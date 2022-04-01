package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type RoleLogic struct {
}

func (RoleLogic) DoAdd(p *models.AddRoleParams) (err error) {
	// 判断是否已存在该role类型
	err = mysql.IsRoleExist(p.Title)
	if err != nil {
		return ErrorRoleExist
	}
	// 增加新role
	err = mysql.AddRole(p)
	if err != nil {
		return ErrorAddRole
	}
	return nil
}

func (RoleLogic) GetRoleList() (roleList []*models.Role, err error) {
	return mysql.GetRoleList()
}
