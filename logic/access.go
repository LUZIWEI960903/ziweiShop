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
