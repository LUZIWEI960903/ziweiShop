package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
)

type ManagerLogic struct {
}

func (ManagerLogic) GetRoleList() (newRoleList []models.NewRoleList, err error) {
	// 查询roleList
	roleList, err := mysql.GetRoleList()
	if err != nil {
		return nil, err
	}
	newRoleList = make([]models.NewRoleList, 0)
	for _, role := range roleList {
		newRole := models.NewRoleList{
			Id:    role.Id,
			Title: role.Title,
		}
		newRoleList = append(newRoleList, newRole)
	}
	return newRoleList, nil
}
