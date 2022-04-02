package models

type Role struct {
	Status      int    // 角色状态
	AddTime     int    // 角色创建时间
	Id          int    // 角色id
	IsDeleted   int    // 是否逻辑删除
	Title       string // 角色名
	Description string // 角色详情
}

func (Role) TableName() string {
	return "role"
}
