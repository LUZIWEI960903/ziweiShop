package shop

import (
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	BaseController
}

// Category 指定商品分类下的商品页面的接口
func (con ProductController) Category(c *gin.Context) {
	id := c.Param("id")
	data, err := logic.ProductLogic{}.SearchProductsById(id)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Category(c *gin.Context)] [logic.ProductLogic{}.SearchProductsById(id)] failed, ", zap.Error(err))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}
