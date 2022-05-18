package models

type User struct {
	Id        int    `json:"id"`         // 用户id
	AddTime   int    `json:"add_time"`   // 创建时间
	Status    int    `json:"status"`     // 状态
	IsDeleted int    `json:"is_deleted"` // 是否逻辑删除
	Phone     string `json:"phone"`      // 用户手机号码
	Password  string `json:"password"`   // 用户密码
	LastIp    string `json:"last_ip"`    // 最后登录的ip
	Email     string `json:"email"`      // 电子邮箱
}

func (User) TableName() string {
	return "user"
}
