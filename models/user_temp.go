package models

type UserTemp struct {
	Id        int    `json:"id"`         // 用户临时表id
	SendCount int    `json:"send_count"` // 发送验证码的次数
	AddTime   int    `json:"add_time"`   // 创建时间
	Ip        string `json:"ip"`         // 用户ip地址
	Email     string `json:"email"`      // 用户手机号码
	AddDay    string `json:"add_day"`    // 手机号增加的日期
	Sign      string `json:"sign"`       // 签名
}

func (UserTemp) TableName() string {
	return "user_temp"
}
