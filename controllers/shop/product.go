package shop

import (
	"strconv"
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductController struct {
	BaseController
}

// Category 指定商品分类下的商品页面的接口
func (con ProductController) Category(c *gin.Context) {
	// 解析参数
	cateId, err1 := strconv.Atoi(c.Param("id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Category(c *gin.Context)] [strconv.Atoi(c.Param(\"id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}
	// 解析分页参数
	page, pageSize, err2 := con.getPagePageSize(c)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Category(c *gin.Context)] [page,pageSize, err2 := con.getPagePageSize(c)] failed, ", zap.Error(err2))
		con.error(c, CodeInValidParams)
		return
	}
	data, err3 := logic.ProductLogic{}.SearchProductsById(cateId, page, pageSize)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Category(c *gin.Context)] [logic.ProductLogic{}.SearchProductsById(cateId, page, pageSize)] failed, ", zap.Error(err3))
		con.error(c, CodeGetDataErr)
		return
	}

	con.success(c, data)
}
