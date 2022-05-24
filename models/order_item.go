package models

type OrderItem struct {
	Id           int     `json:"id,omitempty"`            // 订单详情id
	OrderId      int     `json:"order_id,omitempty"`      // 关联order表的id
	Uid          int     `json:"uid,omitempty"`           // 用户id
	ProductId    int     `json:"product_id,omitempty"`    // 商品id
	ProductNum   int     `json:"product_num,omitempty"`   // 商品数量
	AddTime      int     `json:"add_time,omitempty"`      // 创建数据
	IsDeleted    int     `json:"is_deleted"`              // 是否逻辑删除
	ProductTitle string  `json:"product_title,omitempty"` // 商品名字
	GoodsVersion string  `json:"goods_version,omitempty"` // 商品版本
	GoodsColor   string  `json:"goods_color,omitempty"`   // 商品颜色
	ProductPrice float64 `json:"product_price,omitempty"` // 商品价格
}

func (OrderItem) TableName() string {
	return "order_item"
}
