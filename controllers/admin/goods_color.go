package admin

import (
	"strconv"
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
	// 业务逻辑
	data, err := logic.GoodsColorLogic{}.ShowIndexPageLogic()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsColorController) Index(c *gin.Context)] [logic.GoodsColorLogic{}.ShowIndexPageLogic()] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
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
	// 解析参数
	goodsColorId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsColorController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	//业务逻辑
	goodsColorInfo, err := logic.GoodsColorLogic{}.ShowEditPageLogic(goodsColorId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsColorController) Edit(c *gin.Context)] [logic.GoodsColorLogic{}.ShowEditPageLogic(goodsColorId)] failed, ", zap.Error(err))
		con.error(c, CodeGetEditPageDataErr)
		return
	}

	con.success(c, goodsColorInfo)
}

// DoEdit 编辑商品颜色的接口
func (con GoodsColorController) DoEdit(c *gin.Context) {

}

// Delete 删除商品颜色的接口
func (con GoodsColorController) Delete(c *gin.Context) {

}
