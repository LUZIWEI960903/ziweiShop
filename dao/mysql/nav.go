package mysql

import (
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
func GetNavList() (oNavList []models.Nav, err error) {
	return oNavList, db.Where("is_deleted=0").Find(&oNavList).Error
}
