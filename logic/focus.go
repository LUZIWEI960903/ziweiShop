package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type FocusLogic struct {
}

func (FocusLogic) DoAdd(p *models.AddFocusParams, focusImgSrc string) (err error) {
	return mysql.AddFocus(p, focusImgSrc)
}

func (FocusLogic) GetFocusList() (focusList []models.ResponseFocus, err error) {
	// 获取 focusList
	ofocusList, err := mysql.GetFocusList()
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	for _, focus := range ofocusList {
		newFocus := models.ResponseFocus{
			Id:        focus.Id,
			FocusType: focus.FocusType,
			Sort:      focus.Sort,
			Status:    focus.Status,
			Title:     focus.Title,
			FocusImg:  focus.FocusImg,
			Link:      focus.Link,
		}
		focusList = append(focusList, newFocus)
	}
	return focusList, nil
}
