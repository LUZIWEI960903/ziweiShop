package models

type Manager struct {
	Id       int    // 管理员id
	Status   int    // 管理员状态
	RoleId   int    // 管理员所属角色（部门）
	AddTime  int    // 管理员创建时间
	IsSupper int    // 是否超级管理员
	Username string // 管理员姓名
	Password string // 管理员密码
	Mobile   string // 管理员手机
	Email    string // 管理员邮箱
}

func (Manager) TableName() string {
	return "manager"
}
