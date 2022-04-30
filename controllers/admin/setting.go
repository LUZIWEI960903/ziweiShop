package admin

import (
	"ziweiShop/logic"

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
