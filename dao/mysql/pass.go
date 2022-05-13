package mysql

import "ziweiShop/models"

// IsEmailExist 查询 email是否已存在  --- user 表
func IsEmailExist(email string) bool {
	userList := []models.User{}
	db.Where("is_deleted=0 AND email=?", email).First(&userList)
	if len(userList) > 0 {
		return true
	}
	return false
}
