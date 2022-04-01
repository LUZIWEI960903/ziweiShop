package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// IsRoleExist 查询当前role是否存在
func IsRoleExist(title string) (err error) {
	role := models.Role{}
	RowsAffected := db.Where("title=?", title).First(&role).RowsAffected
	if RowsAffected == 1 {
		return ErrRoleExist
	}
	return nil
}

// AddRole 增加当前role
func AddRole(p *models.AddRoleParams) (err error) {
	role := models.Role{
		Status:      1,
		AddTime:     int(tools.GetUnix()),
		Title:       p.Title,
		Description: p.Description,
	}
	return db.Create(&role).Error
}

// GetRoleList 获取roleList
func GetRoleList() (roleList []*models.Role, err error) {
	roleList = []*models.Role{}
	db.Find(&roleList)
	if len(roleList) < 1 {
		return nil, ErrNoRole
	}
	return roleList, nil
}
