package mysql

import "ziweiShop/models"

// GetTopAccessList 获取顶级模块列表
func GetTopAccessList() (topAccessList []models.Access, err error) {
	topAccessList = []models.Access{}
	err = db.Where("module_id=0 AND is_deleted=0").Find(&topAccessList).Error
	if err != nil {
		return nil, ErrGetTopAccessList
	}
	return topAccessList, nil
}

// AddAccess 添加权限
func AddAccess(access *models.Access) (err error) {
	return db.Create(&access).Error
}

// GetTopAccessListWithAccessList 获取顶级模块及其子权限模块
func GetTopAccessListWithAccessList() (topAccessListWithAccessList []models.Access, err error) {
	topAccessListWithAccessList = []models.Access{}
	err = db.Where("module_id=0 AND is_deleted=0").Preload("AccessList").Find(&topAccessListWithAccessList).Error
	if err != nil {
		return nil, ErrGetTopAccessListWithAccessList
	}
	return topAccessListWithAccessList, nil
}
