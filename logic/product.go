package logic

import (
	"ziweiShop/models"
)

type ProductLogic struct {
	BaseLogic
}

func (l ProductLogic) SearchProductsById(id string) (*models.SearchProductsByKeywordData, error) {
	// 获取基础数据
	baseData, err := l.getBaseData()
	if err != nil {
		return nil, err
	}

	return &models.SearchProductsByKeywordData{
		ShopBaseData: baseData,
	}, nil
}
