package mysql

import (
	"ziweiShop/models"
	"ziweiShop/pkg/tools"
)

// GetAddressByUid 根据用户id 获取 地址列表 --- address 表
func GetAddressByUid(uid int) (addressList []models.Address, err error) {
	return addressList, db.Where("is_deleted=0 AND uid=?", uid).Order("id DESC").Find(&addressList).Error
}

// UpdateDefaultAddress 把当前用户下的所有 default_address 设置为0  --- address 表
func UpdateDefaultAddress(uid int) error {
	return db.Table("address").Where("is_deleted=0 AND uid=?", uid).Updates(map[string]interface{}{"default_address": 0}).Error
}

// CreateAddress 创建收货地址  --- address 表
func CreateAddress(uid int, p *models.AddAddressParams) error {
	address := models.Address{
		Uid:            uid,
		Name:           p.Name,
		DefaultAddress: 1,
		AddTime:        int(tools.GetUnix()),
		Phone:          p.Phone,
		Address:        p.Address,
	}
	return db.Create(&address).Error
}
