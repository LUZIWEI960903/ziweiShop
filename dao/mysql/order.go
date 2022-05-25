package mysql

import "ziweiShop/models"

// CreateOrder  创建order   --- order 表
func CreateOrder(orderObj *models.Order) (*models.Order, error) {
	return orderObj, db.Create(&orderObj).Error
}
