package shop

import (
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

	err := logic.BuyLogic{}.DoCheckout(sign, c)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con BuyController) DoCheckout(c *gin.Context)] [logic.BuyLogic{}.DoCheckout(c)] failed, ", zap.Error(err))
		con.error(c, CodeServerBusy)
		return
	}
	con.success(c, true)
}
