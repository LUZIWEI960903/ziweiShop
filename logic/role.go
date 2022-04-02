package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
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

func (RoleLogic) GetRoleById(roleId int) (data interface{}, err error) {
	// 根据id去查询role信息
	roleInfo, err := mysql.GetRoleById(roleId)
	if err != nil {
		return nil, ErrGetRole
	}
	return gin.H{
		"title":       roleInfo.Title,
		"description": roleInfo.Description,
	}, nil
}
