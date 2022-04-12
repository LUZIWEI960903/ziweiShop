package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// AddGoodsCate 增加商品分类 --- goods_cate 表
func AddGoodsCate(p *models.AddGoodsCateParams, cateImg string) (err error) {
	goodsCate := models.GoodsCate{
		Pid:         p.Pid,
		Status:      p.Status,
		Sort:        p.Sort,
		AddTime:     int(tools.GetUnix()),
		IsDeleted:   0,
		Title:       p.Title,
		CateImg:     cateImg,
		Link:        p.Link,
		Template:    p.Template,
		SubTitle:    p.SubTitle,
		Keywords:    p.Keywords,
		Description: p.Description,
	}
	return db.Create(&goodsCate).Error
}
