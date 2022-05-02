package admin

import (
	"strconv"
	logic "ziweiShop/logic/admin"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

// Index 显示商品类型属性页面的接口
func (con GoodsTypeAttributeController) Index(c *gin.Context) {
	// 获取cate_id
	cateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Index(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err := logic.GoodsTypeAttributeLogic{}.ShowIndexPageLogic(cateId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Index(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.ShowIndexPageLogic(cateId)] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
}

// Add 添加商品类型属性页面的接口
func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	// 获取cate_id
	cateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Add(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err := logic.GoodsTypeAttributeLogic{}.ShowAddPageLogic(cateId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Add(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.ShowAddPageLogic(cateId)] failed, ", zap.Error(err))
		con.error(c, CodeGetAddPageDataErr)
		return
	}

	con.success(c, data)
}

// DoAdd 添加商品类型属性的接口
func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddGoodsTypeAttributeParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.GoodsTypeAttributeLogic{}.DoAddGoodsTypeAttributeLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) DoAdd(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.DoAddGoodsTypeAttributeLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoAddLogicErr)
		return
	}

	con.success(c, true)
}

// Edit 修改商品类型属性页面的接口
func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	// 获取 goodsTypeAttributeId
	goodsTypeAttributeId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err := logic.GoodsTypeAttributeLogic{}.ShowEditPageLogic(goodsTypeAttributeId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Edit(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.ShowEditPageLogic(cateId)] failed, ", zap.Error(err))
		con.error(c, CodeGetEditPageDataErr)
		return
	}

	con.success(c, data)
}

// DoEdit 修改商品类型属性的接口
func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditGoodsTypeAttributeParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.GoodsTypeAttributeLogic{}.DoEditGoodsTypeAttributeLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) DoEdit(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.DoEditGoodsTypeAttributeLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoEditLogicErr)
		return
	}

	con.success(c, true)
}

// Delete 删除商品类型属性的接口
func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	// 解析参数
	goodsTypeAttributeId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Delete(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err = logic.GoodsTypeAttributeLogic{}.DeleteLogic(goodsTypeAttributeId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Delete(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.DeleteLogic(goodsTypeAttributeId)] failed, ", zap.Error(err))
		con.error(c, CodeDeleteLogicErr)
		return
	}

	con.success(c, true)
}
