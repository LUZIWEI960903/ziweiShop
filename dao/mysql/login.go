package mysql

import (
	"fmt"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// IsManagerExist 查询管理员账号是否存在
func IsManagerExist(username string) (err error) {
	u := []models.Manager{}
	db.Where("username=?", username).Find(&u)
	if len(u) < 1 {
		return ErrManagerNotExist
	}
	return nil
}

// VerifyUsernameAndPassword 校验username + password
func VerifyUsernameAndPassword(username, password string) (err error) {
	u := []models.Manager{}
	// md5加密密码
	password = tools.MD5(password)
	db.Where("username=? AND password=?", username, password).Find(&u)
	if len(u) < 1 {
		return ErrManagerPassword
	}
	fmt.Println(u)
	return nil
}
