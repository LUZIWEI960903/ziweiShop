package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type ProductLogic struct {
	BaseLogic
}

func (l ProductLogic) SearchProductsById(cateId, page, pageSize int) (*models.SearchProductsByKeywordData, error) {
	// 获取基础数据
	baseData, err1 := l.getBaseData()
	if err1 != nil {
		return nil, err1
	}

	// 根据cateId 获取商品
	goodsList, pageNum, err2 := mysql.GetGoodsByCateId(cateId, page, pageSize)
	if err2 != nil {
		return nil, err2
	}

	return &models.SearchProductsByKeywordData{
		ShopBaseData: baseData,
		GoodsList:    goodsList,
		PageNum:      pageNum,
		CurrentPage:  page,
	}, nil
}
