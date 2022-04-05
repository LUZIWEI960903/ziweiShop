package logic

import (
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
