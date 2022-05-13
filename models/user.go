package models

type User struct {
	Id        int    // 用户id
	AddTime   int    // 创建时间
	Status    int    // 状态
	IsDeleted int    // 是否逻辑删除
	Phone     string // 用户手机号码
	Password  string // 用户密码
	LastIp    string // 最后登录的ip
	Email     string // 电子邮箱
}

func (User) TableName() string {
	return "user"
}
