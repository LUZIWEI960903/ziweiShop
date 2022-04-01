package routers

import (
	"ziweiShop/controllers/admin"
	"ziweiShop/middlewares"

	"github.com/gin-gonic/gin"
)

func adminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin", middlewares.InitAdminAuthMiddleware)
	{
		// 后台管理登录页面
		adminRouters.GET("/login", admin.LoginController{}.Login)
		// 生成验证码
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		// 后台管理登录
		adminRouters.POST("/login", admin.LoginController{}.DoLogin)
		// 后台管理登出
		adminRouters.GET("/logout", admin.LoginController{}.Logout)

		// 后台主页面
		adminRouters.GET("/", admin.MainPageController{}.Index)

		// 管理员页面
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		// 添加管理员页面
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		// 编辑管理员页面
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		// 删除管理员页面
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		// 轮播图页面
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		// 添加轮播图页面
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		// 编辑轮播图页面
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		// 删除轮播图页面
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

	}
}
