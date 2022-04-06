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
		RowsAffected := db.Where("title=? AND is_deleted=0", params).First(&role).RowsAffected
		if RowsAffected == 1 {
			return ErrRoleExist
		}
	case int: // 根据id查询role是否存在
		RowsAffected := db.Where("id=? AND is_deleted=0", params).First(&role).RowsAffected
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
		IsDeleted:   0,
	}
	return db.Create(&role).Error
}

// GetRoleList 获取roleList
func GetRoleList() (roleList []*models.Role, err error) {
	roleList = []*models.Role{}
	db.Where("is_deleted=0").Find(&roleList)
	if len(roleList) < 1 {
		return nil, ErrNoRole
	}
	return roleList, nil
}

// GetRoleById 根据id获取role信息
func GetRoleById(roleId int) (roleInfo *models.Role, err error) {
	roleInfo = new(models.Role)
	err = db.Where("id=? AND is_deleted=0", roleId).First(&roleInfo).Error
	if err != nil || roleInfo.Title == "" {
		return nil, err
	}
	return roleInfo, nil
}

// EditRole 修改role信息
func EditRole(p *models.EditRoleParams) (err error) {
	role := models.Role{Id: p.Id}
	if RowsAffected := db.Where("is_deleted=0").Find(&role).RowsAffected; RowsAffected != 1 {
		return ErrNoRole
	}
	role.Title = p.Title
	role.Description = p.Description
	return db.Save(&role).Error
}

// DeleteRoleById 根据role id逻辑删除role
func DeleteRoleById(roleId int) (err error) {
	role := models.Role{Id: roleId}
	if RowsAffected := db.Where("is_deleted=0").First(&role).RowsAffected; RowsAffected != 1 {
		return ErrNoRole
	}
	role.IsDeleted = 1
	return db.Save(&role).Error
}

// GetAccessListByRoleId 根据roleId 查询 所有access --- role_access 表
func GetAccessListByRoleId(roleId int) (accessList []models.RoleAccess, err error) {
	accessList = []models.RoleAccess{}
	return accessList, db.Where("role_id=?", roleId).Find(&accessList).Error
}
