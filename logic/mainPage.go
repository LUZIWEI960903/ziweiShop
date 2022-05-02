package logic

import (
	mysql2 "ziweiShop/dao/mysql"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
)

type MainPageLogic struct {
}

func (MainPageLogic) GetAdminIndexPageInfo(managerId int) (data interface{}, err error) {
	// 查询所有 顶级模块 + 子模块
	TopAccessListWithAccessList, err := mysql2.GetTopAccessListWithAccessList()
	if err != nil {
		return nil, ErrorGetTopAccessListWithAccessList
	}

	// 根据ManagerId 查询 其roleId
	managerInfo, err := mysql2.GetManagerInfoById(managerId)
	if err != nil {
		return nil, err
	}

	// 根据roleId 查询 其拥有的 access
	accessList, err := mysql2.GetAccessListByRoleId(managerInfo.RoleId)
	if err != nil {
		return nil, ErrorGetAccessByRoleId
	}
	accessListMap := make(map[int]int, len(accessList))
	for _, access := range accessList {
		// 1:1
		// 2:2
		accessListMap[access.AccessId] = access.AccessId
	}
	// fmt.Println(accessList)

	// 构造返回数据结构
	dataList := make([]models.ResponseTopAccessItemAuth, 0)
	for _, topAccess := range TopAccessListWithAccessList {
		newTopAccess := models.ResponseTopAccessItemAuth{
			AccessId:   topAccess.Id,
			ModuleName: topAccess.ModuleName,
			//IsCheck:    false,
			//AccessItem: nil,
		}
		// 查看该用户是否存在当前 顶级模块 的授权，如果存在 IsCheck=true
		if _, ok := accessListMap[topAccess.Id]; ok {
			newTopAccess.IsCheck = true
		}
		AccessItem := make([]models.ResponseAccessItemAuth, 0)
		for _, access := range topAccess.AccessList {
			newAccess := models.ResponseAccessItemAuth{
				AccessId:   access.Id,
				ActionName: access.ActionName,
				// IsCheck:    false,
			}
			// 查看该用户是否存在当前 子模块 的授权，如果存在 IsCheck=true
			if _, ok := accessListMap[access.Id]; ok {
				newAccess.IsCheck = true
			}
			AccessItem = append(AccessItem, newAccess)
		}
		newTopAccess.AccessItem = AccessItem
		dataList = append(dataList, newTopAccess)
	}
	return gin.H{
		"manager_id": managerId,
		"is_super":   managerInfo.IsSuper,
		"data_list":  dataList,
	}, nil
}
