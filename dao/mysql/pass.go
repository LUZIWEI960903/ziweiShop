package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"

	"github.com/gin-gonic/gin"
)

// IsEmailExist 查询 email是否已存在  --- user 表
func IsEmailExist(email string) bool {
	userList := []models.User{}
	db.Where("is_deleted=0 AND email=?", email).First(&userList)
	if len(userList) > 0 {
		return true
	}
	return false
}

// IsValidRegisterCount 判断 当前ip在当天注册的邮箱个数是否合法  --- user_temp 表
func IsValidRegisterCount(ip, currentDay string) bool {
	var registerCount int64
	db.Table("user_temp").Where("ip=? AND add_day=?", ip, currentDay).Count(&registerCount)
	return registerCount <= 10
}

// GetSendMsgCount 获取 当前邮箱在当天发送验证码的个数  --- user_temp 表
func GetSendMsgCount(email, currentDay string) int {
	userTempList := []models.UserTemp{}
	db.Where("email=? AND add_day=?", email, currentDay).Find(&userTempList)
	if len(userTempList) > 0 {
		return len(userTempList)
	}
	return 0
}

// CreateEmailTemp 创建 记录数据  --- user_temp 表
func CreateEmailTemp(email, currentDay, ip, sign string) error {
	userTemp := models.UserTemp{
		SendCount: 1,
		AddTime:   int(tools.GetUnix()),
		Ip:        ip,
		Email:     email,
		AddDay:    currentDay,
		Sign:      sign,
	}
	return db.Create(&userTemp).Error
}

// GetUserTempInfo 根据 email，currentDay 获取 userTemp信息 --- user_temp 表
func GetUserTempInfo(email, currentDay string) (models.UserTemp, error) {
	userTemp := models.UserTemp{}
	return userTemp, db.Where("email=? AND add_day=?", email, currentDay).First(&userTemp).Error
}

// UpdateMsgCount 更新发送验证码次数  --- user_temp 表
func UpdateMsgCount(userTemp *models.UserTemp) error {
	db.Where("id=?", userTemp.Id)
	userTemp.SendCount += 1
	userTemp.AddTime = int(tools.GetUnix())
	return db.Save(&userTemp).Error
}

// GetUserTempBySign 根据 sign 获取 userTemp  --- user_temp 表
func GetUserTempBySign(sign string) (*models.UserTemp, bool) {
	userTempList := []models.UserTemp{}
	db.Where("sign=?", sign).Find(&userTempList)
	if len(userTempList) > 0 {
		return &userTempList[0], true
	}
	return nil, false
}

// CreateUser 创建用户   --- user 表
func CreateUser(c *gin.Context, email, password string) (models.User, error) {
	user := models.User{
		AddTime:   int(tools.GetUnix()),
		Status:    1,
		IsDeleted: 0,
		Phone:     "",
		Password:  tools.MD5(password),
		LastIp:    c.ClientIP(),
		Email:     email,
	}
	return user, db.Create(&user).Error
}
