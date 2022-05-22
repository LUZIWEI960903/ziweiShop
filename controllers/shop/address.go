package shop

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type AddressController struct {
	BaseController
}

// AddAddress 增加收货地址的接口
func (con AddressController) AddAddress(c *gin.Context) {
	// 解析参数
	p := new(models.AddAddressParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) AddAddress(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	data, err := logic.AddressLogic{}.AddAddress(c, p)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) AddAddress(c *gin.Context)] [logic.AddressLogic{}.AddAddress(c, p)] failed, ", zap.Error(err))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}

// GetOneAddress 获取要修改的收货地址信息的接口
func (con AddressController) GetOneAddress(c *gin.Context) {
	// 解析参数
	addressId, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) GetOneAddress(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err2 := logic.AddressLogic{}.GetOneAddress(c, addressId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) GetOneAddress(c *gin.Context)] [logic.AddressLogic{}.GetOneAddress(c, addressId)] failed, ", zap.Error(err2))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}

// EditAddress 修改收货地址的接口
func (con AddressController) EditAddress(c *gin.Context) {
	// 解析参数
	p := new(models.EditAddressParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) EditAddress(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err := logic.AddressLogic{}.EditAddress(c, p)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con AddressController) EditAddress(c *gin.Context)] [logic.AddressLogic{}.EditAddress(c, p)] failed, ", zap.Error(err))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}
