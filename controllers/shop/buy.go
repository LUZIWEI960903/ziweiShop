package shop

import (
	"strconv"
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BuyController struct {
	BaseController
}

// Checkout 确认订单信息页面的接口
func (con BuyController) Checkout(c *gin.Context) {
	// 业务逻辑
	data, ok := logic.BuyLogic{}.Checkout(c)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con BuyController) Checkout(c *gin.Context)] [logic.BuyLogic{}.Checkout(c)] failed")
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}

// DoCheckout 提交订单执行结算的接口
func (con BuyController) DoCheckout(c *gin.Context) {
	sign := c.PostForm("order_sign")

	data, err := logic.BuyLogic{}.DoCheckout(sign, c)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con BuyController) DoCheckout(c *gin.Context)] [logic.BuyLogic{}.DoCheckout(c)] failed, ", zap.Error(err))
		con.error(c, CodeGetDataErr)
		return
	}
	con.success(c, data)
}

// Pay 支付页面的接口
func (con BuyController) Pay(c *gin.Context) {
	orderId, err1 := strconv.Atoi(c.Query("order_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con BuyController) Pay(c *gin.Context)] [strconv.Atoi(c.Query(\"order_id\"))] failed,", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	data, err2 := logic.BuyLogic{}.Pay(c, orderId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con BuyController) Pay(c *gin.Context)] [logic.BuyLogic{}.Pay(c, orderId)] failed,", zap.Error(err2))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}
