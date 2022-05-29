package shop

import (
	"strconv"
	"ziweiShop/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type AlipayController struct {
	BaseController
}

// Alipay 支付宝支付的接口
func (con AlipayController) Alipay(c *gin.Context) {
	orderId, err1 := strconv.Atoi(c.Query("order_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con AlipayController) Alipay(c *gin.Context)] [strconv.Atoi(c.Query(\"order_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	data, err2 := logic.AlipayLogic{}.Alipay(c, orderId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con AlipayController) Alipay(c *gin.Context)] [logic.AlipayLogic{}.Alipay(c, orderId)] failed, ", zap.Error(err2))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}

// AlipayNotify 支付宝异步订单的接口
func (con AlipayController) AlipayNotify(c *gin.Context) {
	err := logic.AlipayLogic{}.AlipayNotify(c)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con AlipayController) AlipayNotify(c *gin.Context)] [logic.AlipayLogic{}.AlipayNotify(c)] failed, ", zap.Error(err))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, true)
}

// AlipayReturn 支付宝支付返回的接口
func (con AlipayController) AlipayReturn(c *gin.Context) {
	con.success(c, true)
}
