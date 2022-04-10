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

func (FocusLogic) GetFocusById(focusId int) (focusInfo *models.ResponseEditFocus, err error) {
	// 根据 focusId 查询对应的 focus
	focus, err := mysql.GetFocusById(focusId)
	if err != nil {
		return nil, err
	}
	// 构造返回数据
	focusTypeList := []models.FocusType{
		{TypeId: 1, Type: "web"},
		{TypeId: 2, Type: "app"},
		{TypeId: 3, Type: "小程序"},
	}
	newFocusTypeList := make([]models.FocusType, 0)
	newFocusTypeList1 := make([]models.FocusType, 0)
	newFocusTypeList2 := make([]models.FocusType, 0)
	for _, focusType := range focusTypeList {
		if focusType.TypeId == focus.FocusType {
			newFocusTypeList1 = append(newFocusTypeList1, focusType)
		} else {
			newFocusTypeList2 = append(newFocusTypeList2, focusType)
		}
	}
	newFocusTypeList = append(newFocusTypeList1, newFocusTypeList2...)

	focusInfo = &models.ResponseEditFocus{
		Id:        focus.Id,
		FocusType: newFocusTypeList,
		Sort:      focus.Sort,
		Status:    focus.Status,
		Title:     focus.Title,
		FocusImg:  focus.FocusImg,
		Link:      focus.Link,
	}
	return focusInfo, nil
}
