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
