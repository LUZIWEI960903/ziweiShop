package mysql

import "ziweiShop/models"

// GetSetting 获取 系统设置信息  --- setting 表
func GetSetting() (*models.Setting, error) {
	setting := new(models.Setting)
	return setting, db.First(&setting).Error
}

// EditSetting 修改 系统设置信息  --- setting 表
func EditSetting(p *models.Setting) error {
	setting := models.Setting{Id: p.Id}
	err := db.First(&setting).Error
	if err != nil {
		return err
	}

	setting.OssStatus = p.OssStatus
	setting.SiteTitle = p.SiteTitle
	if p.SiteLogo != "" {
		setting.SiteLogo = p.SiteLogo
	}
	setting.SiteKeywords = p.SiteKeywords
	setting.SiteDescription = p.SiteDescription
	if p.DefaultPic != "" {
		setting.DefaultPic = p.DefaultPic
	}
	setting.SiteIcp = p.SiteIcp
	setting.SiteTel = p.SiteTel
	setting.SearchKeywords = p.SearchKeywords
	setting.TongjiCode = p.TongjiCode
	setting.AppId = p.AppId
	setting.AppSecret = p.AppSecret
	setting.EndPoint = p.EndPoint
	setting.BucketName = p.BucketName
	return db.Save(&setting).Error
}
