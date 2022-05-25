package models

type Order struct {
	Id          int     `json:"id,omitempty"`           // 订单id
	OrderId     string  `json:"order_id,omitempty"`     // 订单号码
	Uid         int     `json:"uid,omitempty"`          // 用户id
	AllPrice    float64 `json:"all_price,omitempty"`    // 总价
	Phone       string  `json:"phone,omitempty"`        // 用户手机
	Name        string  `json:"name,omitempty"`         // 订单名
	Address     string  `json:"address,omitempty"`      // 地址
	PayStatus   int     `json:"pay_status,omitempty"`   // 支付状态： 0 未支付  1 已支付
	PayType     int     `json:"pay_type,omitempty"`     // 支付类型： 0 支付宝  1 微信
	OrderStatus int     `json:"order_status,omitempty"` // 订单状态： 0 已下单  1 已付款  2 已配送  3 发货  4 交易成功  5 退货  6 取消
	AddTime     int     `json:"add_time,omitempty"`     // 创建时间
	IsDeleted   int     `json:"is_deleted,omitempty"`   // 是否逻辑删除
}

func (Order) TableName() string {
	return "order"
}
