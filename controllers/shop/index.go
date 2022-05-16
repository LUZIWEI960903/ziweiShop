package shop

import (
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IndexController struct {
	BaseController
}

// Index 商城主页面的接口
func (con IndexController) Index(c *gin.Context) {
	// 业务逻辑
	data, err := logic.ShopIndexLogic{}.ShowIndexPageDataLogic(c)
	if err != nil {
		zap.L().Error("[pkg: shop] [func: (con IndexController) Index(c *gin.Context)] [logic.ShopIndexLogic{}.ShowIndexPageDataLogic(c)] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
}
