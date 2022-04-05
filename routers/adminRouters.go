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
		adminRouters.GET("/login", admin.LoginController{}.Index)
		// 生成验证码
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		// 后台管理登录
		adminRouters.POST("/login", admin.LoginController{}.DoLogin)
		// 后台管理登出
		adminRouters.GET("/logout", admin.LoginController{}.Logout)

		// 后台主页面
		adminRouters.GET("/", admin.MainPageController{}.Index)

		// 角色页面
		adminRouters.GET("/role", admin.RoleController{}.Index)
		// 添加角色页面
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		// 添加角色
		adminRouters.POST("/role/add", admin.RoleController{}.DoAdd)
		// 编辑角色页面
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		// 编辑角色
		adminRouters.PUT("/role/edit", admin.RoleController{}.DoEdit)
		// 删除角色
		adminRouters.DELETE("/role/delete", admin.RoleController{}.Delete)

		// 管理员页面
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		// 添加管理员页面
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		// 添加管理员
		adminRouters.POST("/manager/add", admin.ManagerController{}.DoAdd)
		// 编辑管理员页面
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		// 编辑管理员
		adminRouters.PUT("/manager/edit", admin.ManagerController{}.DoEdit)
		// 删除管理员
		adminRouters.DELETE("/manager/delete", admin.ManagerController{}.Delete)

		// 权限列表页面
		adminRouters.GET("/access", admin.AccessController{}.Index)
		// 添加权限页面
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		// 添加权限
		adminRouters.POST("/access/add", admin.AccessController{}.DoAdd)

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
