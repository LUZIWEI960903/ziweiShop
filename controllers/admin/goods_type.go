package admin

import (
	"strconv"
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
	// 业务逻辑
	goodsTypeList, err := logic.GoodsTypeLogic{}.GetGoodsTypeList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) Index(c *gin.Context)] [logic.GoodsTypeLogic{}.GetGoodsTypeList()] failed, ", zap.Error(err))
		con.error(c, CodeGetGoodsTypeListErr)
		return
	}
	con.success(c, goodsTypeList)
}

// Add 增加商品类型页面的接口
func (con BaseController) Add(c *gin.Context) {
	con.success(c, true)
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
	// 解析参数
	goodsTypeId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	goodsTypeInfo, err := logic.GoodsTypeLogic{}.GetGoodsTypeInfo(goodsTypeId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) Edit(c *gin.Context)] [logic.GoodsTypeLogic{}.GetGoodsTypeInfo(goodsTypeId)] failed, ", zap.Error(err))
		con.error(c, CodeGetGoodsTypeInfoErr)
		return
	}

	con.success(c, goodsTypeInfo)
}

// DoEdit 编辑商品类型的接口
func (con BaseController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditGoodsTypeParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.GoodsTypeLogic{}.DoEdit(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) DoEdit(c *gin.Context)] [logic.GoodsTypeLogic{}.DoEdit(p)] failed, ", zap.Error(err))
		con.error(c, CodeEditGoodsTypeErr)
		return
	}

	con.success(c, true)
}

// Delete 删除商品类型的接口
func (con BaseController) Delete(c *gin.Context) {
	// 解析参数
	goodsTypeId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) Delete(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err = logic.GoodsTypeLogic{}.Delete(goodsTypeId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con BaseController) Delete(c *gin.Context)] [logic.GoodsTypeLogic{}.Delete(goodsTypeId)] failed, ", zap.Error(err))
		con.error(c, CodeDeleteGoodsTypeErr)
		return
	}

	con.success(c, true)
}
