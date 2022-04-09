package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddFocus 增加轮播图  --- focus 表
func AddFocus(p *models.AddFocusParams, focusImgSrc string) (err error) {
	focus := models.Focus{
		FocusType: p.FocusType,
		Sort:      p.Sort,
		Status:    1,
		AddTime:   int(tools.GetUnix()),
		IsDeleted: 0,
		Title:     p.Title,
		FocusImg:  focusImgSrc,
		Link:      p.Link,
	}
	return db.Create(&focus).Error
}
