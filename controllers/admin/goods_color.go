package admin

import (
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsColorController struct {
	BaseController
}

// Index 商品颜色页面的接口
func (con GoodsColorController) Index(c *gin.Context) {

}

// Add 增加商品颜色页面的接口
func (con GoodsColorController) Add(c *gin.Context) {
	con.success(c, true)
}

// DoAdd 增加商品颜色的接口
func (con GoodsColorController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddGoodsColorParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsColorController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.GoodsColorLogic{}.DoAddGoodsColorLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsColorController) DoAdd(c *gin.Context)] [logic.GoodsColorLogic{}.DoAddGoodsColorLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoAddLogicErr)
		return
	}

	con.success(c, true)
}

// Edit 编辑商品颜色页面的接口
func (con GoodsColorController) Edit(c *gin.Context) {

}

// DoEdit 编辑商品颜色的接口
func (con GoodsColorController) DoEdit(c *gin.Context) {

}

// Delete 删除商品颜色的接口
func (con GoodsColorController) Delete(c *gin.Context) {

}
