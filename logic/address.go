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
