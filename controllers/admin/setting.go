package admin

import (
	"ziweiShop/logic"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SettingController struct {
	BaseController
}

// Index 系统管理页面的接口
func (con SettingController) Index(c *gin.Context) {
	// 业务逻辑
	data, err := logic.SettingLogic{}.ShowIndexPageDataLogic()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con SettingController) Index(c *gin.Context)] [logic.SettingLogic{}.ShowIndexPageDataLogic()] failed, ", zap.Error(err))
		con.error(c, CodeGetIndexPageDataErr)
		return
	}

	con.success(c, data)
}

// DoEdit 修改系统管理的信息的接口
func (con SettingController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.Setting)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con SettingController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err := logic.SettingLogic{}.EditSettingLogic(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con SettingController) DoEdit(c *gin.Context)] [logic.SettingLogic{}.EditSettingLogic(p)] failed, ", zap.Error(err))
		con.error(c, CodeDoEditLogicErr)
		return
	}

	con.success(c, true)
}
