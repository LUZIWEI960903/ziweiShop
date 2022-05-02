package logic

import (
	mysql "ziweiShop/dao/mysql/admin"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
)

type RoleLogic struct {
}

func (RoleLogic) DoAdd(p *models.AddRoleParams) (err error) {
	// 判断是否已存在该role类型
	err = mysql.IsRoleExist(p.Title)
	if err != nil {
		return ErrorRoleExist
	}
	// 增加新role
	err = mysql.AddRole(p)
	if err != nil {
		return ErrorAddRole
	}
	return nil
}

func (RoleLogic) GetRoleList() (roleList []*models.Role, err error) {
	return mysql.GetRoleList()
}

func (RoleLogic) GetRoleById(roleId int) (data interface{}, err error) {
	// 根据id去查询role信息
	roleInfo, err := mysql.GetRoleById(roleId)
	if err != nil {
		return nil, ErrorGetRole
	}
	return gin.H{
		"title":       roleInfo.Title,
		"description": roleInfo.Description,
	}, nil
}

func (RoleLogic) DoEdit(p *models.EditRoleParams) (err error) {
	// 判断是否已存在该role类型
	err = mysql.IsRoleExist(p.Title)
	if err != nil {
		return ErrorRoleExist
	}
	// 修改role title，description
	err = mysql.EditRole(p)
	if err != nil {
		return ErrorEditRole
	}
	return nil
}

func (RoleLogic) DeleteRoleById(roleId int) (err error) {
	return mysql.DeleteRoleById(roleId)
}

func (RoleLogic) Auth(roleId int) (data []models.ResponseTopAccessItemAuth, err error) {
	// 查询该roleId是否存在
	_, err = mysql.GetRoleById(roleId)
	if err != nil {
		return nil, ErrorRoleNotExist
	}

	// 获取所有的access (所有的 顶级模块+其子模块)
	TopAccessListWithAccessList, err := mysql.GetTopAccessListWithAccessList()
	if err != nil {
		return nil, ErrorGetTopAccessListWithAccessList
	}

	// 根据roleId 查询其所有access
	accessList, err := mysql.GetAccessListByRoleId(roleId)
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
		data = append(data, newTopAccess)
	}
	return data, nil
}

func (RoleLogic) DoAuth(p *models.DoAuthParams) (err error) {
	// 校验 roleId是否存在
	_, err = mysql.GetRoleById(p.Id)
	if err != nil {
		return ErrorRoleNotExist
	}

	// 校验 顶级模块 + 子模块 是否存在
	accessIdList := make([]int, 0)
	for _, topAccessList := range p.AccessItem {
		if _, err := mysql.GetAccessById(topAccessList.AccessId); err != nil {
			return err
		}
		accessIdList = append(accessIdList, topAccessList.AccessId)
		for _, accessList := range topAccessList.AccessItem {
			if _, err := mysql.GetAccessById(accessList.AccessId); err != nil {
				return err
			}
			accessIdList = append(accessIdList, accessList.AccessId)
		}
	}

	// 清空 该 roleId 下的所有 access
	err = mysql.DeleteAccessByRoleId(p.Id)
	if err != nil {
		return ErrorDeleteAccessByRoleId
	}

	// 根据accessId 插入数据
	return mysql.CreateRoleAccess(p.Id, accessIdList)
}
