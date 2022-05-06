package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type ShopIndexLogic struct {
	BaseLogic
}

func (l ShopIndexLogic) ShowIndexPageDataLogic() (*models.ShopIndexPageData, error) {
	// 获取基础数据
	baseData, err := l.getBaseData()
	if err != nil {
		return nil, err
	}

	// 获取轮播图数据
	oType1FocusList, err2 := mysql.GetType1FocusList()
	if err2 != nil {
		return nil, err2
	}

	// 获取cateId=1的商品数据
	cateId1Goods := mysql.GetGoodsByCategory(1, "hot", 2)

	// 获取cateId=xxx
	// ...
	// ...

	// 构造返回数据
	// --轮播图
	FocusList := make([]models.ResponseFocus, 0)
	for _, oFocus := range oType1FocusList {
		newFocus := models.ResponseFocus{
			Id:        oFocus.Id,
			FocusType: oFocus.FocusType,
			Sort:      oFocus.Sort,
			Status:    oFocus.Status,
			Title:     oFocus.Title,
			FocusImg:  oFocus.FocusImg,
			Link:      oFocus.Link,
		}
		FocusList = append(FocusList, newFocus)
	}

	return &models.ShopIndexPageData{
		ShopBaseData:     baseData,
		FocusList:        FocusList,
		CateId1GoodsList: cateId1Goods,
	}, nil
}
