package mysql

import "ziweiShop/models"

// CreateOrder  创建order   --- order 表
func CreateOrder(orderObj *models.Order) (*models.Order, error) {
	return orderObj, db.Create(&orderObj).Error
}

// GetOrderById  查询 order  --- order 表
func GetOrderById(orderId int) (*models.Order, error) {
	order := models.Order{}
	return &order, db.Where("is_deleted=0 AND id=?", orderId).Select("uid,order_id,phone,name,address").First(&order).Error
}
