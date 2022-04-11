package mysql

import (
	"ziweiShop/models"

	"gorm.io/gorm"
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
	err = db.Where("module_id=0 AND is_deleted=0").Order("sort DESC").Preload("AccessList", "is_deleted=0", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Find(&topAccessListWithAccessList).Error
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

// IsTopAccess 根据accessId 判断是否为顶级模块
func IsTopAccess(accessId int) bool {
	access := models.Access{}
	if RowsAffected := db.Where("id=? AND is_deleted=0 AND module_id=0", accessId).First(&access).RowsAffected; RowsAffected != 1 {
		return false
	}
	return true
}

// GetSonAccessList 查询 module_id = accessId
func GetSonAccessList(accessId int) bool {
	access := []models.Access{}
	db.Where("module_id=? AND is_deleted=0", accessId).Find(&access)
	if len(access) > 0 {
		// 存在子模块
		return true
	}
	return false
}

// DeleteAccessById 根据 accessId 逻辑删除 access
func DeleteAccessById(accessId int) (err error) {
	access := models.Access{}
	if RowsAffected := db.Where("id=? AND is_deleted=0", accessId).First(&access).RowsAffected; RowsAffected != 1 {
		return ErrGetAccess
	}
	access.IsDeleted = 1
	return db.Save(&access).Error
}

// GetAccessIdByUrl 根据url 查询accessId  --- access 表
func GetAccessIdByUrl(url string) (accessId int) {
	access := []models.Access{}
	db.Where("url=? AND is_deleted=0", url).First(&access)
	if len(access) < 1 {
		return -1
	}
	return access[0].Id
}

// GetAllAccessUrl 获取所有url  --- access 表
func GetAllAccessUrl() (accessUrlList []models.ResponseAccessUrl) {
	accessList := []models.Access{}
	db.Find(&accessList)
	if len(accessList) < 1 {
		return nil
	}
	for _, access := range accessList {
		accessUrl := models.ResponseAccessUrl{Url: access.Url}
		accessUrlList = append(accessUrlList, accessUrl)
	}
	return accessUrlList
}
