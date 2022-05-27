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

// GetAddressById  根据address id查询 address信息   --- address 表
func GetAddressById(addressId int) (address *models.Address, err error) {
	return address, db.Where("is_deleted=0 AND id=?", addressId).Select("id,uid,name,phone,address,default_address").First(&address).Error
}

// UpdateAddress  修改address 信息  --- address 表
func UpdateAddress(p *models.EditAddressParams) (address *models.Address, err error) {
	err1 := db.Where("is_deleted=0 AND id=?", p.Id).First(&address).Error
	if err1 != nil || address == nil {
		return nil, err1
	}
	address.Address = p.Address
	address.Phone = p.Phone
	address.Name = p.Name
	return address, db.Save(&address).Error
}

// UpdateDefaultAddress1  将所有 该用户的 address 设为非默认地址 并把 当前 address 设置为默认地址  --- address 表
func UpdateDefaultAddress1(uid, addressId int) ([]models.Address, error) {
	err1 := db.Table("address").Where("is_deleted=0 AND uid=?", uid).Updates(map[string]interface{}{"default_address": 0}).Error
	if err1 != nil {
		return nil, err1
	}
	address := models.Address{}
	err2 := db.Where("is_deleted=0 AND id=?", addressId).First(&address).Error
	if err2 != nil {
		return nil, err2
	}
	address.DefaultAddress = 1
	err3 := db.Save(&address).Error
	if err3 != nil {
		return nil, err3
	}
	addressList := []models.Address{}
	return addressList, db.Where("is_deleted=0 AND uid=?", uid).Find(&addressList).Error
}

// GetDefaultAddressByUid  根据uid 获取默认地址  --- address 表
func GetDefaultAddressByUid(uid int) (address *models.Address, err error) {
	addressList := []models.Address{}
	err = db.Where("is_deleted=0 AND uid=? AND default_address=1", uid).First(&addressList).Error
	if err != nil || len(addressList) < 1 {
		return nil, ErrGetData
	}
	return &addressList[0], nil
}
