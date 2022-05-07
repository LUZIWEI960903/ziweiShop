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

// Detail 商品详情页面的接口
func (con ProductController) Detail(c *gin.Context) {
	// 解析参数
	goodsId, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Detail(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err2 := logic.ProductLogic{}.GetGoodsInfoData(goodsId)
	if err2 != nil {
		zap.L().Error("[pkg: shop] [func: (con ProductController) Detail(c *gin.Context)] [logic.ProductLogic{}.GetGoodsInfoData(goodsId)] failed, ", zap.Error(err2))
		con.error(c, CodeGetDataErr)
		return
	}
	con.success(c, data)
}
