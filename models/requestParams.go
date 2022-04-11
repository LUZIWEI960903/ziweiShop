package models

type LoginParams struct {
	Username     string `json:"username"`      // 用户名
	Password     string `json:"password"`      // 用户密码
	CaptchaId    string `json:"captcha_id"`    // 验证码id
	CaptchaValue string `json:"captcha_value"` // 验证码的值
}

type AddRoleParams struct {
	Title       string `json:"title"`       // 角色名
	Description string `json:"description"` // 角色详情
}

type EditRoleParams struct {
	Id          int    `json:"id"`          // role的id
	Title       string `json:"title"`       // 角色名
	Description string `json:"description"` // 角色详情
}

type AddManagerParams struct {
	RoleId   int    `json:"role_id"`  // 管理员所属角色（部门）
	Username string `json:"username"` // 管理员姓名
	Password string `json:"password"` // 管理员密码
	Mobile   string `json:"mobile"`   // 管理员手机
	Email    string `json:"email"`    // 管理员邮箱
}

type EditManagerParams struct {
	Id       int    `json:"id"`       // 管路员id
	RoleId   int    `json:"role_id"`  // 管理员所属角色（部门）
	Password string `json:"password"` // 管理员密码
	Mobile   string `json:"mobile"`   // 管理员手机
	Email    string `json:"email"`    // 管理员邮箱
}

type AddAccessParams struct {
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	Status      int    `json:"status"`      // 模块状态
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type EditAccessParams struct {
	Id          int    `json:"id"`          // 权限id
	Type        int    `json:"type"`        // 模块类型: 1.表示模块 2.表示菜单 3.表示操作
	ModuleId    int    `json:"module_id"`   // 与当前Id自关联
	Sort        int    `json:"sort"`        // 排序
	Status      int    `json:"status"`      // 模块状态
	ActionName  string `json:"action_name"` // 操作名称
	ModuleName  string `json:"module_name"` // 模块名称
	Url         string `json:"url"`         // 路由跳转地址
	Description string `json:"description"` // 描述
}

type DoAuthParams struct {
	Id         int                         `json:"role_id"`     // 角色id
	AccessItem []ResponseTopAccessItemAuth `json:"access_item"` // 权限列表
}

type AddFocusParams struct {
	FocusType int    `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    `json:"sort"`       // 排序方式
	Title     string `json:"title"`      // 轮播图名称
	Link      string `json:"links"`      // 跳转地址
}

type EditFocusParams struct {
	Id        int    `json:"id"`         // 轮播图id
	FocusType int    `json:"focus_type"` // 轮播图类型 1.web 2.app 3.小程序
	Sort      int    `json:"sort"`       // 排序方式
	Status    int    `json:"status"`     // 状态
	Title     string `json:"title"`      // 轮播图名称
	Link      string `json:"link"`       // 跳转地址
}
