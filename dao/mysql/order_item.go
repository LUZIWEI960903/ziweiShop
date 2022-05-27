package mysql

import "ziweiShop/models"

// CreateOrderItem  创建 orderItem   --- order_item 表
func CreateOrderItem(orderItem *models.OrderItem) {
	db.Create(&orderItem)
}

// GetOrderItemByOrderId  查询order item  --- order_item 表
func GetOrderItemByOrderId(orderId int) ([]models.OrderItem, error) {
	orderItemList := []models.OrderItem{}
	err := db.Where("is_deleted=0 AND order_id=?", orderId).Find(&orderItemList).Error
	return orderItemList, err
}
