package models

type NewRoleList struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerListRole struct {
	Id    int    `json:"id"`    // role的id
	Title string `json:"title"` // 角色名
}

type IndexManagerList struct {
	Id       int                  `json:"id"`       // manager的id
	Username string               `json:"username"` // manager的名字
	Mobile   string               `json:"mobile"`   // manager的电话
	Email    string               `json:"email"`    // manager的邮箱
	AddTime  string               `json:"add_time"` // manager的创建时间
	Role     IndexManagerListRole `json:"role"`     // manager对应role的信息
}

type EditManagerList struct {
	Id       int           `json:"id"`        // manager的id
	Username string        `json:"username"`  // manager的名字
	Password string        `json:"password"`  // manager的密码
	Mobile   string        `json:"mobile"`    // manager的电话
	Email    string        `json:"email"`     // manager的邮箱
	RoleList []NewRoleList `json:"role_list"` // 所有role的信息
}

type NewTopAccess struct {
	Id         int    `json:"id"`          // 权限id
	ModuleName string `json:"module_name"` // 模块名称
}

type NewAccessList struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type NewTopAccessListWithAccessList struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述

	AccessList []NewAccessList `json:"access_list"` // 该权限下的模块
}

type AccessType struct {
	Id       int    `json:"id"`        // access type id
	TypeName string `json:"type_name"` // 模块类型名字: 1.表示模块 2.表示菜单 3.表示操作
}

type EditAccessInfo struct {
	Id          int          `json:"id"`          // 权限id
	TypeList    []AccessType `json:"type_list"`   // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int          `json:"module_id"`   // 与当前Id自关联
	Sort        int          `json:"sort"`        // 排序
	ActionName  string       `json:"action_name"` // 操作名称
	ModuleName  string       `json:"module_name"` // 模块名称
	Url         string       `json:"url"`         // 路由跳转地址
	Description string       `json:"description"` // 描述

	NewTopAccessList []NewTopAccess `json:"access_list"` // 所有的顶级模块
}

type ResponseAccessItemAuth struct {
	AccessId   int    `json:"access_id"`   // 权限id
	ActionName string `json:"action_name"` // 操作名称
	IsCheck    bool   `json:"is_check"`    // 是否已授权
}

type ResponseTopAccessItemAuth struct {
	AccessId   int                      `json:"access_id"`       // 权限id
	ModuleName string                   `json:"module_name"`     // 模块名字
	IsCheck    bool                     `json:"is_check"`        // 是否已授权
	AccessItem []ResponseAccessItemAuth `json:"son_access_list"` // 子模块列表
}
