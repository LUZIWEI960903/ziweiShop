package mysql

import (
	"ziweiShop/models"
)

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

// GetAccessById 根据accessId 获取 access 信息
func GetAccessById(accessId int) (accessInfo *models.Access, err error) {
	accessInfo = &models.Access{}
	err = db.Where("id=? AND is_deleted=0", accessId).First(&accessInfo).Error
	if err != nil {
		return nil, ErrGetAccess
	}
	return accessInfo, nil
}

// EditAccess 修改 access 信息
func EditAccess(p *models.EditAccessParams) (err error) {
	// 查询access 是否存在
	access := models.Access{}
	if RowsAffected := db.Where("is_deleted=0 AND id=?", p.Id).Find(&access).RowsAffected; RowsAffected != 1 {
		return ErrGetAccess
	}
	// 修改 信息
	access.Type = p.Type
	access.ModuleId = p.ModuleId
	access.Sort = p.Sort
	access.Status = p.Status
	access.ActionName = p.ActionName
	access.ModuleName = p.ModuleName
	access.Url = p.Url
	access.Description = p.Description
	return db.Save(&access).Error
}
