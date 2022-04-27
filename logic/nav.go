package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type NavLogic struct {
}

func (NavLogic) AddNavLogic(p *models.AddNavParams) error {
	return mysql.AddNav(p)
}

func (NavLogic) ShowIndexPageDataLogic() (*models.NavIndexPageData, error) {
	// 获取 所有nav
	oNavList, err := mysql.GetNavList()
	if err != nil {
		return nil, err
	}

	// 构造返回数据
	NavItems := make([]models.NavList, 0)
	for _, oNav := range oNavList {
		newNav := models.NavList{
			Id:       oNav.Id,
			Position: oNav.Position,
			Sort:     oNav.Sort,
			Status:   oNav.Status,
			Title:    oNav.Title,
			Link:     oNav.Link,
			Relation: oNav.Relation,
		}
		NavItems = append(NavItems, newNav)
	}

	return &models.NavIndexPageData{
		NavItems: NavItems,
	}, nil
}
