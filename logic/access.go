package logic

import (
	"time"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type AccessLogic struct {
}

func (AccessLogic) GetTopAccessList() (newTopAccessList []models.NewTopAccess, err error) {
	// 获取完整topAccessList
	topAccessList, err := mysql.GetTopAccessList()
	if err != nil {
		return nil, ErrorGetTopAccessList
	}
	// 构造返回的topAccessList
	for _, topAccess := range topAccessList {
		newTopAccess := models.NewTopAccess{
			Id:         topAccess.Id,
			ModuleName: topAccess.ModuleName,
		}
		newTopAccessList = append(newTopAccessList, newTopAccess)
	}
	return newTopAccessList, nil
}

func (AccessLogic) DoAdd(p *models.AddAccessParams) (err error) {
	// 构造module
	access := &models.Access{
		Type:        p.Type,
		ModuleId:    p.ModuleId,
		Sort:        p.Sort,
		AddTime:     int(time.Now().Unix()),
		Status:      p.Status,
		IsDeleted:   0,
		ActionName:  p.ActionName,
		ModuleName:  p.ModuleName,
		Url:         p.Url,
		Description: p.Description,
	}
	return mysql.AddAccess(access)
}

func (AccessLogic) GetTopAccessListWithAccessList() (newTopAccessListWithAccessList []models.NewTopAccessListWithAccessList, err error) {
	// 查询出完整的 topAccessListWithAccessList
	topAccessListWithAccessList, err := mysql.GetTopAccessListWithAccessList()
	if err != nil {
		return nil, ErrorGetTopAccessListWithAccessList
	}
	// 构造 newTopAccessListWithAccessList
	for _, TopAccessList := range topAccessListWithAccessList {
		newTopAccessList := models.NewTopAccessListWithAccessList{
			Id:          TopAccessList.Id,
			Type:        TopAccessList.Type,
			ModuleId:    TopAccessList.ModuleId,
			Sort:        TopAccessList.Sort,
			ActionName:  TopAccessList.ActionName,
			ModuleName:  TopAccessList.ModuleName,
			Url:         TopAccessList.Url,
			Description: TopAccessList.Description,
		}
		allAccessList := make([]models.NewAccessList, 0)
		for _, accessList := range TopAccessList.AccessList {
			newAccessList := models.NewAccessList{
				Id:          accessList.Id,
				Type:        accessList.Type,
				ModuleId:    accessList.ModuleId,
				Sort:        accessList.Sort,
				ActionName:  accessList.ActionName,
				ModuleName:  accessList.ModuleName,
				Url:         accessList.Url,
				Description: accessList.Description,
			}
			allAccessList = append(allAccessList, newAccessList)
		}
		newTopAccessList.AccessList = allAccessList
		newTopAccessListWithAccessList = append(newTopAccessListWithAccessList, newTopAccessList)
	}
	return newTopAccessListWithAccessList, nil
}

func (AccessLogic) GetAccessById(accessId int) (editAccessInfo *models.EditAccessInfo, err error) {
	// 根据 accessId 查询 access
	accessInfo, err := mysql.GetAccessById(accessId)
	if err != nil {
		return nil, ErrorGetAccessInfo
	}
	// 查询所有顶级模块
	topAccessList, err := mysql.GetTopAccessList()
	if err != nil {
		return nil, ErrorGetTopAccessList
	}
	// 构造 newTopAccessList
	newTopAccessList := make([]models.NewTopAccess, 0)
	newTopAccessList1 := make([]models.NewTopAccess, 0)
	newTopAccessList2 := make([]models.NewTopAccess, 0)
	for _, topAccess := range topAccessList {
		newTopAccess := models.NewTopAccess{
			Id:         topAccess.Id,
			ModuleName: topAccess.ModuleName,
		}
		if topAccess.Id == accessInfo.ModuleId {
			newTopAccessList1 = append(newTopAccessList1, newTopAccess)
		} else {
			newTopAccessList2 = append(newTopAccessList2, newTopAccess)
		}
	}
	newTopAccessList = append(newTopAccessList1, newTopAccessList2...)
	// 构造 newType
	newType := make([]models.AccessType, 0)
	newType1 := make([]models.AccessType, 0)
	newType2 := make([]models.AccessType, 0)
	typeList := []models.AccessType{
		{Id: 1, TypeName: "模块"},
		{Id: 2, TypeName: "菜单"},
		{Id: 3, TypeName: "操作"},
	}
	for _, t := range typeList {
		if accessInfo.Type == t.Id {
			newType1 = append(newType1, t)
		} else {
			newType2 = append(newType2, t)
		}
	}
	newType = append(newType1, newType2...)
	// 构造返回数据
	editAccessInfo = &models.EditAccessInfo{
		Id:               accessInfo.Id,
		TypeList:         newType,
		ModuleId:         accessInfo.ModuleId,
		Sort:             accessInfo.Sort,
		ActionName:       accessInfo.ActionName,
		ModuleName:       accessInfo.ModuleName,
		Url:              accessInfo.Url,
		Description:      accessInfo.Description,
		NewTopAccessList: newTopAccessList,
	}
	return editAccessInfo, nil
}

func (AccessLogic) DoEdit(p *models.EditAccessParams) (err error) {
	return mysql.EditAccess(p)
}
