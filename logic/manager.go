package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

type ManagerLogic struct {
}

func (ManagerLogic) GetRoleList() (newRoleList []models.NewRoleList, err error) {
	// 查询roleList
	roleList, err := mysql.GetRoleList()
	if err != nil {
		return nil, err
	}
	//newRoleList = make([]models.NewRoleList, 0)
	for _, role := range roleList {
		newRole := models.NewRoleList{
			Id:    role.Id,
			Title: role.Title,
		}
		newRoleList = append(newRoleList, newRole)
	}
	return newRoleList, nil
}

func (ManagerLogic) DoAdd(p *models.AddManagerParams) (err error) {
	// 根据username查询管理员是否已存在，存在则报错
	err = mysql.GetManagerByUsername(p.Username)
	if err != nil {
		return ErrorManagerExist
	}
	// 判断role id 是否存在
	err = mysql.IsRoleExist(p.RoleId)
	if err != mysql.ErrRoleExist {
		return ErrorRoleNotExist
	}
	// 增加管理员
	return mysql.AddManager(p)
}

func (ManagerLogic) GetIndexManagerList() (indexManagerList []models.IndexManagerList, err error) {
	// 获取完整managerList
	managerList, err := mysql.GetManagerList()
	if err != nil {
		return nil, err
	}
	// 重组managerList到IndexManagerList
	for _, manager := range managerList {
		newManager := models.IndexManagerList{
			Id:       manager.Id,
			Username: manager.Username,
			Mobile:   manager.Mobile,
			Email:    manager.Email,
			AddTime:  tools.UnixToDate(manager.AddTime),
			Role: models.IndexManagerListRole{
				Id:    manager.RoleId,
				Title: manager.Role.Title,
			},
		}
		indexManagerList = append(indexManagerList, newManager)
	}
	return indexManagerList, nil
}
