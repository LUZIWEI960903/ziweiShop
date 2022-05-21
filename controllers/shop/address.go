package shop

import (
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
