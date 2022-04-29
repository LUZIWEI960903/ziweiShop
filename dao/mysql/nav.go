package mysql

import (
	"math"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddNav 增加 nav  --- nav 表
func AddNav(p *models.AddNavParams) error {
	nav := models.Nav{
		Position:  p.Position,
		IsOpennew: p.IsOpennew,
		Sort:      p.Sort,
		Status:    p.Status,
		AddTime:   int(tools.GetUnix()),
		Title:     p.Title,
		Link:      p.Link,
		Relation:  p.Relation,
	}

	return db.Create(&nav).Error
}

// GetNavList 获取 所有 nav  --- nav 表
func GetNavList(page, pageSize int) (oNavList []models.Nav, pageCount float64, err error) {
	err = db.Where("is_deleted=0").Offset((page - 1) * pageSize).Limit(pageSize).Find(&oNavList).Error

	var count int64
	db.Table("nav").Count(&count)

	pageCount = math.Ceil(float64(count) / float64(pageSize))
	return oNavList, pageCount, err
}

// GetNavById  根据 navId 查询 信息  --- nav 表
func GetNavById(navId int) (*models.Nav, error) {
	oNavList := []models.Nav{}
	err := db.Where("id=? AND is_deleted=0", navId).First(&oNavList).Error
	if len(oNavList) < 1 || err != nil {
		return nil, ErrGetNav
	}
	return &oNavList[0], nil
}

// EditNav 修改 nav 信息  --- nav 表
func EditNav(p *models.EditNavParams) error {
	oNavList := []models.Nav{}
	err := db.Where("id=? AND is_deleted=0", p.Id).First(&oNavList).Error
	if len(oNavList) < 1 || err != nil {
		return ErrGetNav
	}

	oNavList[0].Id = p.Id
	oNavList[0].Position = p.Position
	oNavList[0].IsOpennew = p.IsOpennew
	oNavList[0].Sort = p.Sort
	oNavList[0].Status = p.Status
	oNavList[0].Title = p.Title
	oNavList[0].Link = p.Link
	oNavList[0].Relation = p.Relation
	return db.Save(&oNavList).Error
}

// DeleteNav 根据 navId  逻辑删除 nav  --- nav 表
func DeleteNav(navId int) error {
	oNavList := []models.Nav{}
	err := db.Where("id=? AND is_deleted=0", navId).First(&oNavList).Error
	if len(oNavList) < 1 || err != nil {
		return ErrGetNav
	}

	oNavList[0].IsDeleted = 1
	return db.Save(&oNavList).Error
}
