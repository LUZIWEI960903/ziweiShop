package mysql

import "ziweiShop/models"

// GetSetting 获取 系统设置信息  --- setting 表
func GetSetting() (*models.Setting, error) {
	setting := new(models.Setting)
	return setting, db.First(&setting).Error
}
