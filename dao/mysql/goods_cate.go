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

// GetTopGoodsCateList 查询所有 顶级商品分类  --- goods_cate 表
func GetTopGoodsCateList() (oTopGoodsCateList []models.GoodsCate, err error) {
	oTopGoodsCateList = []models.GoodsCate{}
	err = db.Where("pid=0 AND is_deleted=0").Find(&oTopGoodsCateList).Error
	return oTopGoodsCateList, err
}

// GetTopGoodsCateWithGoodsCateList 查询所有 顶级商品分类+子商品分类  --- goods_cate 表
func GetTopGoodsCateWithGoodsCateList() (oTopGoodsCateWithGoodsCateList []models.GoodsCate, err error) {
	oTopGoodsCateWithGoodsCateList = []models.GoodsCate{}
	err = db.Where("pid=0 AND is_deleted=0").Preload("GoodsCateItems", "is_deleted=0").Find(&oTopGoodsCateWithGoodsCateList).Error
	return oTopGoodsCateWithGoodsCateList, err
}

// GetGoodsCateById 根据 goodsCateId 查询该商品分类信息  --- goods_cate 表
func GetGoodsCateById(goodsCateId int) (ogoodsCateInfo *models.GoodsCate, err error) {
	ogoodsCateList := []models.GoodsCate{}
	err = db.Where("id=? AND is_deleted=0", goodsCateId).First(&ogoodsCateList).Error
	if len(ogoodsCateList) < 1 {
		return nil, ErrGetGoodsCate
	}
	return &ogoodsCateList[0], err
}
