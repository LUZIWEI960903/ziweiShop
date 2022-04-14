package admin

import (
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsTypeController struct {
	BaseController
}

// Index 商品类型页面的接口
func (con BaseController) Index(c *gin.Context) {

}

// Add 增加商品类型页面的接口
func (con BaseController) Add(c *gin.Context) {

}

// DoAdd 增加商品类型的接口
func (con BaseController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddGoodsTypeParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.GoodsTypeLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) DoAdd(c *gin.Context)] [logic.GoodsTypeLogic{}.DoAdd(p)] failed, ", zap.Error(err))
		con.error(c, CodeAddGoodsTypeErr)
		return
	}

	con.success(c, true)
}

// Edit 编辑商品类型页面的接口
func (con BaseController) Edit(c *gin.Context) {

}

// DoEdit 编辑商品类型的接口
func (con BaseController) DoEdit(c *gin.Context) {

}

// Delete 删除商品类型的接口
func (con BaseController) Delete(c *gin.Context) {

}
