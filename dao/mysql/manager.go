package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// GetManagerByUsername 判断管理员是否存在
func GetManagerByUsername(username string) (err error) {
	u := []models.Manager{}
	db.Where("username=?", username).Find(&u)
	if len(u) > 0 {
		return ErrManagerExist
	}
	return nil
}

// AddManager 增加管理员
func AddManager(p *models.AddManagerParams) (err error) {
	manager := models.Manager{
		Status:   1,
		RoleId:   p.RoleId,
		AddTime:  int(tools.GetUnix()),
		IsSuper:  0,
		Username: p.Username,
		Password: tools.MD5(p.Password),
		Mobile:   p.Mobile,
		Email:    p.Email,
	}
	return db.Create(&manager).Error
}
