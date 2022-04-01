package mysql

import (
	"errors"
	"fmt"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// IsManagerExist 查询管理员账号是否存在
func IsManagerExist(username string) (err error) {
	u := []models.Manager{}
	db.Where("username=?", username).Find(&u)
	if len(u) < 1 {
		return errors.New("Manager is not exist~~")
	}
	return nil
}

// VerifyUsernameAndPassword
func VerifyUsernameAndPassword(username, password string) (err error) {
	u := []models.Manager{}
	// md5加密密码
	password = tools.MD5(password)
	db.Where("username=? AND password=?", username, password).Find(&u)
	if len(u) < 1 {
		return errors.New("Manager username or password failed~~")
	}
	fmt.Println(u)
	return nil
}
