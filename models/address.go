package models

type Address struct {
	Id             int    `json:"id"`              // 地址id
	Uid            int    `json:"uid"`             // 用户id
	Name           int    `json:"name"`            // 地址名
	Phone          string `json:"phone"`           // 手机号码
	Address        string `json:"address"`         // 详情地址
	DefaultAddress int    `json:"default_address"` // 是否默认地址
	AddTime        int    `json:"add_time"`        // 创建时间
	IsDeleted      int    `json:"is_deleted"`      // 是否逻辑删除
}

func (Address) TableName() string {
	return "address"
}
