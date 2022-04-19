package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type GoodsLogic struct {
}

func (GoodsLogic) ShowAddPageLogic() (data *models.GoodsAddPageData, err error) {
	// 获取商品分类
	oTopGoodsCateWithGoodsCateList, err := mysql.GetTopGoodsCateWithGoodsCateList()
	if err != nil {
		return nil, err
	}
	// 获取商品颜色信息

	// 获取商品规格包装

	// 构造返回数据
	GoodCateItems := make([]models.GoodsCateWithGoodsCate1, 0)
	for _, TopGoodsCate := range oTopGoodsCateWithGoodsCateList {
		newTopGoodsCate := models.GoodsCateWithGoodsCate1{
			Id:    TopGoodsCate.Id,
			Pid:   TopGoodsCate.Pid,
			Title: TopGoodsCate.Title,
			//ChildGoodsCateItems: nil,
		}
		ChildGoodsCateItems := make([]models.TopGoodsCate2, 0)
		for _, GoodsCate := range TopGoodsCate.GoodsCateItems {
			newGoodsCate := models.TopGoodsCate2{
				Id:    GoodsCate.Id,
				Pid:   GoodsCate.Pid,
				Title: GoodsCate.Title,
			}
			ChildGoodsCateItems = append(ChildGoodsCateItems, newGoodsCate)
		}
		newTopGoodsCate.ChildGoodsCateItems = ChildGoodsCateItems
		GoodCateItems = append(GoodCateItems, newTopGoodsCate)
	}

	return &models.GoodsAddPageData{
		GoodsCateItems: GoodCateItems,
	}, nil
}
