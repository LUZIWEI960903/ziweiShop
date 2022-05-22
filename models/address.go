package models

type Address struct {
	Id             int    `json:"id,omitempty"`              // 地址id
	Uid            int    `json:"uid,omitempty"`             // 用户id
	DefaultAddress int    `json:"default_address,omitempty"` // 是否默认地址
	AddTime        int    `json:"add_time,omitempty"`        // 创建时间
	IsDeleted      int    `json:"is_deleted,omitempty"`      // 是否逻辑删除
	Name           string `json:"name,omitempty"`            // 地址名
	Phone          string `json:"phone,omitempty"`           // 手机号码
	Address        string `json:"address,omitempty"`         // 详情地址
}

func (Address) TableName() string {
	return "address"
}
