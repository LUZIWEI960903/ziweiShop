package admin

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type GoodsTypeAttributeController struct {
	BaseController
}

// Index 显示商品类型属性页面的接口
func (con GoodsTypeAttributeController) Index(c *gin.Context) {

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
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) Add(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err))
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
	err := logic.GoodsTypeAttributeLogic{}.DoAddLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con GoodsTypeAttributeController) DoAdd(c *gin.Context)] [logic.GoodsTypeAttributeLogic{}.DoAddLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoAddLogicErr)
		return
	}

	con.success(c, true)
}
