package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// IsRoleExist 查询当前role是否存在
func IsRoleExist(params interface{}) (err error) {
	role := models.Role{}
	switch params.(type) {
	case string: // 根据title查询role是否存在
		RowsAffected := db.Where("title=?", params).First(&role).RowsAffected
		if RowsAffected == 1 {
			return ErrRoleExist
		}
	case int: // 根据id查询role是否存在
		RowsAffected := db.Where("id=?", params).First(&role).RowsAffected
		if RowsAffected == 1 {
			return ErrRoleExist
		}
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

// GetRoleById 根据id获取role信息
func GetRoleById(roleId int) (roleInfo *models.Role, err error) {
	roleInfo = new(models.Role)
	err = db.Where("id=?", roleId).First(&roleInfo).Error
	if err != nil || roleInfo.Title == "" {
		return nil, err
	}
	return roleInfo, nil
}
