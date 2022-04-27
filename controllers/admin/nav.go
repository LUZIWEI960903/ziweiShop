package admin

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type NavController struct {
	BaseController
}

// Index 导航栏主页面的接口
func (con NavController) Index(c *gin.Context) {
	// 业务逻辑
	data, err := logic.NavLogic{}.ShowIndexPageDataLogic()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) Index(c *gin.Context)] [logic.NavLogic{}.ShowIndexPageDataLogic()] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
}

// Add 增加导航栏页面的接口
func (con NavController) Add(c *gin.Context) {
	con.success(c, true)
}

// DoAdd 增加导航栏的接口
func (con NavController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddNavParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	//业务逻辑
	err := logic.NavLogic{}.AddNavLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) DoAdd(c *gin.Context)] [logic.NavLogic{}.AddNavLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoAddLogicErr)
		return
	}

	con.success(c, true)
}

// Edit 编辑导航栏页面的接口
func (con NavController) Edit(c *gin.Context) {
	// 解析参数
	navId, err1 := strconv.Atoi(c.Query("id"))
	if err1 != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) Edit(c *gin.Context)] [strconv.Atoi(c.Query(\"id\"))] failed, ", zap.Error(err1))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	data, err2 := logic.NavLogic{}.ShowEditPageLogic(navId)
	if err2 != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) Edit(c *gin.Context)] [logic.NavLogic{}.ShowEditPageLogic(navId)] failed, ", zap.Error(err2))
		con.error(c, CodeGetEditPageDataErr)
		return
	}

	con.success(c, data)
}

// DoEdit 编辑导航栏的接口
func (con NavController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditNavParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	err := logic.NavLogic{}.EditNavLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con NavController) DoEdit(c *gin.Context)] [logic.NavLogic{}.EditNavLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoEditLogicErr)
		return
	}

	con.success(c, true)
}

// Delete 删除导航栏的接口
func (con NavController) Delete(c *gin.Context) {

}
