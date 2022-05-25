package mysql

import "ziweiShop/models"

// CreateOrderItem  创建 orderItem   --- order_item 表
func CreateOrderItem(orderItem *models.OrderItem) {
	db.Create(&orderItem)
}
