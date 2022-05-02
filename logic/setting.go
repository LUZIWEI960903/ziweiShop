package logic

import (
	"ziweiShop/dao/mysql"
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
