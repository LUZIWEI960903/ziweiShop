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

func (GoodsCateLogic) GetTopGoodsCateWithGoodsCateList() (TopGoodsCateWithGoodsCateList []models.GoodsCateWithGoodsCate, err error) {
	// 查看 顶级商品分类 + 子商品分类
	oTopGoodsCateWithGoodsCateList, err := mysql.GetTopGoodsCateWithGoodsCateList()
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	for _, topGoodsCate := range oTopGoodsCateWithGoodsCateList {
		newTopGoodsCate := models.GoodsCateWithGoodsCate{
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
			// GoodsCateItems: nil,
		}
		for _, goodsCate := range topGoodsCate.GoodsCateItems {
			newGoodsCate := models.TopGoodsCate{
				Id:          goodsCate.Id,
				Pid:         goodsCate.Pid,
				Status:      goodsCate.Status,
				Sort:        goodsCate.Sort,
				Title:       goodsCate.Title,
				CateImg:     goodsCate.CateImg,
				Link:        goodsCate.Link,
				Template:    goodsCate.Template,
				SubTitle:    goodsCate.SubTitle,
				Keywords:    goodsCate.Keywords,
				Description: goodsCate.Description,
			}
			newTopGoodsCate.GoodsCateItems = append(newTopGoodsCate.GoodsCateItems, newGoodsCate)
		}
		TopGoodsCateWithGoodsCateList = append(TopGoodsCateWithGoodsCateList, newTopGoodsCate)
	}
	return TopGoodsCateWithGoodsCateList, nil
}

func (GoodsCateLogic) GetGoodsCateInfo(goodsCateId int) (goodsCateInfo *models.GoodsCateInfo, err error) {
	// 根据 goodsCateId 查询该商品分类信息
	ogoodsCateInfo, err := mysql.GetGoodsCateById(goodsCateId)
	if err != nil {
		return nil, err
	}
	// 查询所有 顶级商品分类
	otopGoodsCateList, err := mysql.GetTopGoodsCateList()
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	topGoodsCateList := make([]models.TopGoodsCate1, 0)
	topGoodsCateList1 := make([]models.TopGoodsCate1, 0)
	topGoodsCateList2 := make([]models.TopGoodsCate1, 0)
	for _, topGoodsCate := range otopGoodsCateList {
		newTopGoodsCate := models.TopGoodsCate1{
			Id:    topGoodsCate.Id,
			Pid:   topGoodsCate.Pid,
			Title: topGoodsCate.Title,
		}
		if topGoodsCate.Id == ogoodsCateInfo.Pid {
			topGoodsCateList1 = append(topGoodsCateList1, newTopGoodsCate)
		} else {
			topGoodsCateList2 = append(topGoodsCateList2, newTopGoodsCate)
		}
		topGoodsCateList = append(topGoodsCateList1, topGoodsCateList2...)
	}
	goodsCateInfo = &models.GoodsCateInfo{
		Id:            ogoodsCateInfo.Id,
		Status:        ogoodsCateInfo.Status,
		Sort:          ogoodsCateInfo.Sort,
		Title:         ogoodsCateInfo.Title,
		CateImg:       ogoodsCateInfo.CateImg,
		Link:          ogoodsCateInfo.Link,
		Template:      ogoodsCateInfo.Template,
		SubTitle:      ogoodsCateInfo.SubTitle,
		Keywords:      ogoodsCateInfo.Keywords,
		Description:   ogoodsCateInfo.Description,
		TopGoodsCate1: topGoodsCateList,
	}

	return goodsCateInfo, nil
}
