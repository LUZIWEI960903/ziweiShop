package models

type RoleAccess struct {
	RoleId   int // 角色id
	AccessId int // 权限id
}

func (RoleAccess) TableName() string {
	return "role_access"
}
