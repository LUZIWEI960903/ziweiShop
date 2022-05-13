package models

type UserTemp struct {
	Id        int    // 用户临时表id
	SendCount int    // 发送验证码的次数
	AddTime   int    // 创建时间
	Ip        string // 用户ip地址
	Email     string // 用户手机号码
	AddDay    string // 手机号增加的日期
	Sign      string // 签名
}

func (UserTemp) TableName() string {
	return "user_temp"
}
