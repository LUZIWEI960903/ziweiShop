package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

	"github.com/gin-gonic/gin"
)

type AddressLogic struct {
}

func (AddressLogic) AddAddress(c *gin.Context, p *models.AddAddressParams) (*models.AddAddressData, error) {
	// 获取用户信息
	user := models.User{}
	cookie.Cookie.Get(c, "userInfo", &user)

	// 根据用户id获取地址列表,地址超过10则不被允许
	addressList, err1 := mysql.GetAddressByUid(user.Id)
	if err1 != nil || len(addressList) > 10 {
		return nil, err1
	}

	// 更新所有默认地址=0
	err2 := mysql.UpdateDefaultAddress(user.Id)
	if err2 != nil {
		return nil, err2
	}

	// 增加收货地址，默认地址=0
	err3 := mysql.CreateAddress(user.Id, p)
	if err3 != nil {
		return nil, err3
	}

	// 获取所有收货地址
	newAddressList, err4 := mysql.GetAddressByUid(user.Id)
	if err4 != nil {
		return nil, err4
	}

	return &models.AddAddressData{
		AddressList: newAddressList,
	}, nil
}

func (AddressLogic) GetOneAddress(c *gin.Context, addressId int) (*models.GetOneAddressData, error) {
	// 获取用户信息
	user := models.User{}
	cookie.Cookie.Get(c, "userInfo", &user)

	// 根据 address id查询
	address, err1 := mysql.GetAddressById(addressId)
	if err1 != nil {
		return nil, err1
	}

	// 校验address uid 与 user 的id
	if user.Id != address.Uid {
		return nil, ErrorInvalidParams
	}

	return &models.GetOneAddressData{
		AddressInfo: *address,
	}, nil
}

func (AddressLogic) EditAddress(c *gin.Context, p *models.EditAddressParams) (*models.EditAddressData, error) {
	// 获取用户信息
	user := models.User{}
	cookie.Cookie.Get(c, "userInfo", &user)

	// 根据 address id查询
	address, err1 := mysql.GetAddressById(p.Id)
	if err1 != nil {
		return nil, err1
	}

	// 校验address uid 与 user 的id
	if user.Id != address.Uid {
		return nil, ErrorInvalidParams
	}

	// 修改 address 信息
	_, err2 := mysql.UpdateAddress(p)
	if err2 != nil {
		return nil, err2
	}

	// 查询 address 列表
	addressList, err3 := mysql.GetAddressByUid(user.Id)
	if err3 != nil {
		return nil, err3
	}

	return &models.EditAddressData{
		AddressList: addressList,
	}, nil
}
