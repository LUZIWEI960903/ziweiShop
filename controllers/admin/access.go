package admin

import (
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

// Index 显示权限列表页面的接口
func (con AccessController) Index(c *gin.Context) {

}

// Add 显示增加权限页面的接口
func (con AccessController) Add(c *gin.Context) {
	// 业务逻辑
	topAccessList, err := logic.AccessLogic{}.GetTopAccessList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Add(c *gin.Context)] [logic.AccessLogic{}.GetTopAccessList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetTopAccessListErr)
		return
	}
	con.success(c, topAccessList)
}

// DoAdd 执行增加权限的接口
func (con AccessController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddAccessParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 参数校验
	if err := verifyAddAccessParams(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [verifyAddAccessParams(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEmptyModuleName)
		return
	}
	// 业务逻辑
	err := logic.AccessLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [logic.AccessLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		con.error(c, CodeAddAccessErr)
		return
	}
	con.success(c, true)
}
