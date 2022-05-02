package logic

import (
	mysql "ziweiShop/dao/mysql/admin"
	"ziweiShop/models"
)

type SettingLogic struct {
}

func (SettingLogic) ShowIndexPageDataLogic() (*models.Setting, error) {
	return mysql.GetSetting()
}

func (SettingLogic) EditSettingLogic(p *models.Setting) error {
	return mysql.EditSetting(p)
}
