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
