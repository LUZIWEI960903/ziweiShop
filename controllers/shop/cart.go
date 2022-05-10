package shop

import (
	"strconv"
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CartController struct {
	BaseController
}

// Get 获取购物车数据 显示购物车数据的接口
func (con CartController) Get(c *gin.Context) {
	// 业务逻辑
	data := logic.CartLogic{}.GetCart(c)
	con.success(c, data)
}

// AddCart 添加购物车的接口
func (con CartController) AddCart(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("goods_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) AddCart(c *gin.Context)] [strconv.Atoi(c.Query(\"goods_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}
	colorId, _ := strconv.Atoi(c.Query("color_id"))

	// 业务逻辑
	err2 := logic.CartLogic{}.AddCart(c, goodsId, colorId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) AddCart(c *gin.Context)] [logic.CartLogic{}.AddCart(c,goodsId, colorId)] failed, ", zap.Error(err2))
		con.error(c, CodeAddCartErr)
		return
	}

	con.success(c, true)
}
