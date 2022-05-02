package routers

import (
	"ziweiShop/controllers/shop"

	"github.com/gin-gonic/gin"
)

func shopRoutersInit(r *gin.Engine) {
	shopRouters := r.Group("/")
	{
		// 商城主页面
		shopRouters.GET("/", shop.IndexController{}.Index)
	}
}
