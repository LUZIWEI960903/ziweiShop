package routers

import (
	"ziweiShop/controllers/admin"

	"github.com/gin-gonic/gin"
)

func adminRoutersInit(r *gin.Engine) {
	adminRouters := r.Group("/admin")
	{
		// 后台主页面
		adminRouters.GET("/", admin.MainPageController{}.Index)
		// 后台欢迎页面
		adminRouters.GET("/welcome", admin.MainPageController{}.Welcome)

		// 后台管理登录页面
		adminRouters.GET("/login", admin.LoginController{}.Index)
		// 后台管理登录
		adminRouters.POST("/login", admin.LoginController{}.DoLogin)

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
