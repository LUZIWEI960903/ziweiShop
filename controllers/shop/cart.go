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

// AddCartSuccess 增加购物车成功的页面
func (con CartController) AddCartSuccess(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) AddCartSuccess(c *gin.Context)] [strconv.Atoi(c.Query(\"goods_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err2 := logic.CartLogic{}.AddCartSuccess(goodsId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) AddCartSuccess(c *gin.Context)] [logic.CartLogic{}.AddCartSuccess(goodsId)] failed, ", zap.Error(err2))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}

// IncCart ajax增加购物车商品数量的接口
func (con CartController) IncCart(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("goods_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) IncCart(c *gin.Context)] [strconv.Atoi(c.Query(\"goods_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}
	goodsColor := c.Query("goods_color")

	// 业务逻辑
	data := logic.CartLogic{}.IncCart(c, goodsId, goodsColor)
	con.success(c, data)
}

// DecCart ajax减少购物车商品数量的接口
func (con CartController) DecCart(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("goods_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) DecCart(c *gin.Context)] [strconv.Atoi(c.Query(\"goods_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	goodsColor := c.Query("goods_color")
	// 业务逻辑
	data := logic.CartLogic{}.DecCart(c, goodsId, goodsColor)
	con.success(c, data)
}

// ChangeOneCart ajax改变一个购物车选择的状态的接口
func (con CartController) ChangeOneCart(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("goods_id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con CartController) ChangeOneCart(c *gin.Context)] [strconv.Atoi(c.Query(\"goods_id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	goodsColor := c.Query("goods_color")
	// 业务逻辑
	data := logic.CartLogic{}.ChangeOneCart(c, goodsId, goodsColor)
	con.success(c, data)
}
