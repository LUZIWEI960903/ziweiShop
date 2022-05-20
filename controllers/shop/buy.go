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
