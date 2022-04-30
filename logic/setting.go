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
