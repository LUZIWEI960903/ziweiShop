package admin

import (
	"net/http"
	"ziweiShop/logic"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con ManagerController) Add(c *gin.Context) {
	// 返回roleList
	var managerService = logic.ManagerLogic{}
	roleList, err := managerService.GetRoleList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Add(c *gin.Context)] [managerService.GetRoleList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetRoleListErr)
		return
	}
	con.success(c, roleList)
}

func (con ManagerController) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con ManagerController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
