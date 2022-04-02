package admin

import (
	"errors"
	"net/http"
	"ziweiShop/logic"
	"ziweiShop/models"

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

func (con ManagerController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddManagerParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err := logic.ManagerLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoAdd(c *gin.Context)] [logic.ManagerLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		if errors.Is(err, logic.ErrorManagerExist) {
			con.error(c, CodeManagerExistErr)
			return
		}
		con.error(c, CodeAddManagerErr)
		return
	}
	con.success(c, true)
}

func (con ManagerController) Edit(c *gin.Context) {

}

func (con ManagerController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
