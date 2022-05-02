package mysql

import (
	"fmt"
	"path"
	"reflect"
	"strconv"
	"strings"
	"ziweiShop/models"

	. "github.com/hunterhug/go_image"
)

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

// GetSettingFromColumn 通过列名来获取系统设置的值  --- setting 表
func GetSettingFromColumn(columnName string) string {
	setting := models.Setting{}
	db.First(&setting)
	// 通过反射来获取
	v := reflect.ValueOf(setting)
	return v.FieldByName(columnName).String()
}

// ResizeGoodsImage 根据setting 的ThumbnailSize 生成商品缩放图  --- setting 表
func ResizeGoodsImage(filename string) {
	// 获取后缀名
	extname := path.Ext(filename)

	thumbnailSize := GetSettingFromColumn("ThumbnailSize")
	thumbnailSizeSlice := strings.Split(thumbnailSize, ",")
	for i, l := 0, len(thumbnailSizeSlice); i < l; i++ {
		w := thumbnailSizeSlice[i]
		wInt, _ := strconv.Atoi(w)
		savepath := "static/upload/resize/" + filename + "_" + w + "x" + w + extname
		err := ThumbnailF2F(filename, savepath, wInt, wInt)
		if err != nil {
			fmt.Printf("生成按宽度高度缩放图:%s\n", err.Error())
		} else {
			fmt.Printf("生成按宽度高度缩放图:%s\n", savepath)
		}
	}
}
