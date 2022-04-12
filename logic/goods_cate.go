package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsCateLogic struct {
}

func (GoodsCateLogic) GetTopGoodsCateList() (TopGoodsCateList []models.TopGoodsCate, err error) {
	// 查询所有 顶级商品分类
	oTopGoodsCateList, err := mysql.GetTopGoodsCateList()
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	for _, topGoodsCate := range oTopGoodsCateList {
		newTopGoodsCate := models.TopGoodsCate{
			Id:          topGoodsCate.Id,
			Pid:         topGoodsCate.Pid,
			Status:      topGoodsCate.Status,
			Sort:        topGoodsCate.Sort,
			Title:       topGoodsCate.Title,
			CateImg:     topGoodsCate.CateImg,
			Link:        topGoodsCate.Link,
			Template:    topGoodsCate.Template,
			SubTitle:    topGoodsCate.SubTitle,
			Keywords:    topGoodsCate.Keywords,
			Description: topGoodsCate.Description,
		}
		TopGoodsCateList = append(TopGoodsCateList, newTopGoodsCate)
	}
	return TopGoodsCateList, nil
}

func (GoodsCateLogic) DoAdd(p *models.AddGoodsCateParams, cateImg string) (err error) {
	return mysql.AddGoodsCate(p, cateImg)
}
