package routers

import (
	"ziweiShop/logger"
	"ziweiShop/middlewares"

	"github.com/gin-gonic/gin"
)

func Init(cfgMode string) *gin.Engine {
	if cfgMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())

	// swagger
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 后台管理相关路由
	adminRoutersInit(r)

	return r
}
