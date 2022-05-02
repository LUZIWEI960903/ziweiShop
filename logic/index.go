package logic

import (
	"ziweiShop/dao/mysql"
)

type ShopIndexLogic struct {
}

func (ShopIndexLogic) ShowIndexPageDataLogic() (*models.ShopIndexPageData, error) {
	// 获取顶部导航
	topNavList, err1 := mysql.GetTopNavList()
	if err1 != nil {
		return nil, err1
	}
}
