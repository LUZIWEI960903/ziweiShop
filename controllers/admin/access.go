package admin

import (
	"ziweiShop/logic"

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
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Add(c *gin.Context)] [managerService.GetRoleList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetTopAccessListErr)
		return
	}
	con.success(c, topAccessList)
}

// DoAdd 执行增加权限的接口
func (con AccessController) DoAdd(c *gin.Context) {

}
