package models

type Access struct {
	Id          int    // 权限id
	Type        int    // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    // 与当前Id自关联
	Sort        int    // 排序
	AddTime     int    // 创建时间
	Status      int    // 模块状态
	IsDeleted   int    // 是否逻辑删除
	ActionName  string // 操作名称
	ModuleName  string // 模块名称
	Url         string // 路由跳转地址
	Description string // 描述

	AccessList []Access `gorm:"foreignKey:ModuleId;references:Id"` // 自关联
}

func (Access) TableName() string {
	return "access"
}
