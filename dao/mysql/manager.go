package mysql

import (
	"strings"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// GetManagerByIdOrUsername 判断管理员是否存在
func GetManagerByIdOrUsername(params interface{}) (err error) {
	u := models.Manager{}
	switch params.(type) {
	case string:
		RowsAffected := db.Where("username=?", params).Find(&u).RowsAffected
		if RowsAffected == 1 {
			return ErrManagerExist
		}
	case int:
		RowsAffected := db.Where("id=?", params).Find(&u).RowsAffected
		if RowsAffected == 1 {
			return ErrManagerExist
		}
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

// GetManagerList 获取完整的managerList
func GetManagerList() (managerList []*models.Manager, err error) {
	managerList = []*models.Manager{}
	db.Preload("Role").Find(&managerList)
	if len(managerList) > 0 {
		return managerList, nil
	}
	return nil, ErrGetManagerList
}

// GetManagerInfoById 根据managerId获取信息
func GetManagerInfoById(managerId int) (managerInfo *models.Manager, err error) {
	managerInfo = &models.Manager{}
	err = db.Where("id=?", managerId).First(&managerInfo).Error
	if err != nil {
		return nil, ErrManagerNotExist
	}
	return managerInfo, nil
}

// EditManager 修改manager信息
func EditManager(p *models.EditManagerParams) (err error) {
	oldManagerInfo := models.Manager{}
	db.Where("id=?", p.Id).First(&oldManagerInfo)
	if password := strings.Trim(p.Password, " "); password != "" {
		// password非空，代表修改了密码
		oldManagerInfo.Password = tools.MD5(p.Password)
	}
	oldManagerInfo.RoleId = p.RoleId
	oldManagerInfo.Mobile = p.Mobile
	oldManagerInfo.Email = p.Email
	return db.Save(&oldManagerInfo).Error
}
