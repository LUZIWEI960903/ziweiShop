package routers

import (
	"ziweiShop/logger"
	"ziweiShop/middlewares"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

const cookieSecret = "github.com/LUZIWEI960903/ziweiShop"

func Init(cfgMode string) *gin.Engine {
	if cfgMode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) // 设置发布模式
	}

	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.Cors())

	// swagger
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 配置cookie、sessions
	store := cookie.NewStore([]byte(cookieSecret))
	r.Use(sessions.Sessions("mysession", store))

	// 后台管理相关路由
	adminRoutersInit(r)

	// 商品前台相关路由
	shopRoutersInit(r)

	return r
}
