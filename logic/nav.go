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

func (NavLogic) ShowEditPageLogic(navId int) (*models.NavEditPageData, error) {
	// 查询 nav
	oNav, err := mysql.GetNavById(navId)
	if err != nil {
		return nil, err
	}

	NavInfo := models.Nav1{
		Id:        oNav.Id,
		Position:  oNav.Position,
		IsOpennew: oNav.IsOpennew,
		Sort:      oNav.Sort,
		Status:    oNav.Status,
		Title:     oNav.Title,
		Link:      oNav.Link,
		Relation:  oNav.Relation,
	}

	return &models.NavEditPageData{
		NavInfo: NavInfo,
	}, nil
}

func (NavLogic) EditNavLogic(p *models.EditNavParams) error {
	return mysql.EditNav(p)
}

func (NavLogic) DeleteNavLogic(navId int) error {
	return mysql.DeleteNav(navId)
}
