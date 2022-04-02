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
