package logic

import (
	mysql2 "ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

type ManagerLogic struct {
}

func (ManagerLogic) GetRoleList() (newRoleList []models.NewRoleList, err error) {
	// 查询roleList
	roleList, err := mysql2.GetRoleList()
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
	err = mysql2.GetManagerByIdOrUsername(p.Username)
	if err != nil {
		return ErrorManagerExist
	}
	// 判断role id 是否存在
	err = mysql2.IsRoleExist(p.RoleId)
	if err != mysql2.ErrRoleExist {
		return ErrorRoleNotExist
	}
	// 增加管理员
	return mysql2.AddManager(p)
}

func (ManagerLogic) GetIndexManagerList() (indexManagerList []models.IndexManagerList, err error) {
	// 获取完整managerList
	managerList, err := mysql2.GetManagerList()
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

func (ManagerLogic) GetEditManagerInfo(managerId int) (editManagerInfo *models.EditManagerList, err error) {
	// 判断是否存在该managerId
	err = mysql2.GetManagerByIdOrUsername(managerId)
	if err != mysql2.ErrManagerExist {
		return nil, ErrorManagerNotExist
	}
	// 根据managerId查询manager信息
	managerInfo, err := mysql2.GetManagerInfoById(managerId)
	if err != nil {
		return nil, err
	}
	// 获取所有role信息
	roleList, err := mysql2.GetRoleList()
	if err != nil {
		return nil, err
	}
	newRoleList := make([]models.NewRoleList, 0)
	newRoleList1 := make([]models.NewRoleList, 0) // 存放当前managerId的role
	newRoleList2 := make([]models.NewRoleList, 0) // 存放当前managerId以外的role
	for _, role := range roleList {
		newRole := models.NewRoleList{
			Id:    role.Id,
			Title: role.Title,
		}
		if newRole.Id == managerInfo.RoleId {
			newRoleList1 = append(newRoleList1, newRole)
		} else {
			newRoleList2 = append(newRoleList2, newRole)
		}
	}
	// 使当前managerId的role排头部
	newRoleList = append(newRoleList1, newRoleList2...)
	// 构造返回数据
	editManagerInfo = &models.EditManagerList{
		Id:       managerInfo.Id,
		Username: managerInfo.Username,
		Password: "",
		Mobile:   managerInfo.Mobile,
		Email:    managerInfo.Email,
		RoleList: newRoleList,
	}
	return editManagerInfo, nil
}

func (ManagerLogic) DoEdit(p *models.EditManagerParams) (err error) {
	// 查询managerId 是否存在
	err = mysql2.GetManagerByIdOrUsername(p.Id)
	if err != mysql2.ErrManagerExist {
		return ErrorManagerNotExist
	}
	// 查询roleId 是否存在
	err = mysql2.IsRoleExist(p.RoleId)
	if err != mysql2.ErrRoleExist {
		return ErrorRoleNotExist
	}
	// 执行修改manager信息操作
	return mysql2.EditManager(p)
}

func (ManagerLogic) DeleteManager(managerId int) (err error) {
	return mysql2.DeleteManagerById(managerId)
}
