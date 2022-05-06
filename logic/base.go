package logic

import (
	"strings"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type BaseLogic struct {
}

func (l BaseLogic) getBaseData() (*models.ShopBaseData, error) {
	// 获取顶部导航
	oTopNavList, err1 := mysql.GetTopNavList()
	if err1 != nil {
		return nil, err1
	}

	// 获取分类的数据
	oTopGoodsCateWithGoodsCateList, err2 := mysql.GetTopGoodsCateWithGoodsCateList()
	if err2 != nil {
		return nil, err2
	}

	// 获取中间导航
	oMiddleNavList, err3 := mysql.GetMiddleNavList()
	if err3 != nil {
		return nil, err3
	}

	// 构造返回数据
	// --顶部导航
	TopNavList := make([]models.TopNav, 0)
	for _, oTopNav := range oTopNavList {
		newTopNav := models.TopNav{
			Id:        oTopNav.Id,
			IsOpennew: oTopNav.IsOpennew,
			Sort:      oTopNav.Sort,
			Status:    oTopNav.Status,
			Title:     oTopNav.Title,
			Link:      oTopNav.Link,
			Relation:  oTopNav.Relation,
		}
		TopNavList = append(TopNavList, newTopNav)
	}

	// --商品分类
	TopGoodsCateWithGoodsCateList := make([]models.GoodsCateWithGoodsCate, 0)
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

	// --中间导航
	MiddleNavList := make([]models.MiddleNav, 0)
	for _, oMiddleNav := range oMiddleNavList {
		newMiddleNav := models.MiddleNav{
			Id:        oMiddleNav.Id,
			IsOpennew: oMiddleNav.IsOpennew,
			Sort:      oMiddleNav.Sort,
			Status:    oMiddleNav.Status,
			Title:     oMiddleNav.Title,
			Link:      oMiddleNav.Link,
			Relation:  oMiddleNav.Relation,
			//GoodsItems:
		}
		// 处理Relation 字符串 21,22,23
		newMiddleNav.Relation = strings.ReplaceAll(newMiddleNav.Relation, "，", ",")
		relationIds := strings.Split(newMiddleNav.Relation, ",")
		GoodsList, _ := mysql.GetGoodsListByIds(relationIds)
		newMiddleNav.GoodsItems = GoodsList

		MiddleNavList = append(MiddleNavList, newMiddleNav)
	}
	return &models.ShopBaseData{
		TopNavList:                 TopNavList,
		GoodsCateWithGoodsCateList: TopGoodsCateWithGoodsCateList,
		MiddleNavList:              MiddleNavList,
	}, nil
}
