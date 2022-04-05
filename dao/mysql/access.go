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
