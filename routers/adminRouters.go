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
		// 给角色授权页面
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		// 给角色授权
		adminRouters.POST("/role/auth", admin.RoleController{}.DoAuth)

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
		// 编辑权限页面
		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		// 编辑权限
		adminRouters.PUT("/access/edit", admin.AccessController{}.DoEdit)
		// 删除权限
		adminRouters.DELETE("/access/delete", admin.AccessController{}.Delete)

		// 轮播图页面
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		// 添加轮播图页面
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		// 添加轮播图
		adminRouters.POST("/focus/add", admin.FocusController{}.DoAdd)
		// 编辑轮播图页面
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		// 编辑轮播图
		adminRouters.PUT("/focus/edit", admin.FocusController{}.DoEdit)
		// 删除轮播图页面
		adminRouters.DELETE("/focus/delete", admin.FocusController{}.Delete)

		// Ajax更新 指定表的 指定id的 status
		adminRouters.PUT("/changeStatus", admin.AJAXController{}.ChangeStatus)
		// Ajax更新 指定表的 指定id的 sort
		adminRouters.PUT("/changeSort", admin.AJAXController{}.ChangeSort)

		// 商品分类页面
		adminRouters.GET("/goodsCate", admin.GoodsCateController{}.Index)
		// 添加商品分类页面
		adminRouters.GET("/goodsCate/add", admin.GoodsCateController{}.Add)
		// 添加商品分类
		adminRouters.POST("/goodsCate/add", admin.GoodsCateController{}.DoAdd)
		// 编辑商品分类页面
		adminRouters.GET("/goodsCate/edit", admin.GoodsCateController{}.Edit)
		// 编辑商品分类
		adminRouters.PUT("/goodsCate/edit", admin.GoodsCateController{}.DoEdit)
		// 删除商品分类页面
		adminRouters.DELETE("/goodsCate/delete", admin.GoodsCateController{}.Delete)

		// 商品类型页面
		adminRouters.GET("/goodsType", admin.GoodsTypeController{}.Index)
		// 添加商品类型页面
		adminRouters.GET("/goodsType/add", admin.GoodsTypeController{}.Add)
		// 添加商品类型
		adminRouters.POST("/goodsType/add", admin.GoodsTypeController{}.DoAdd)
		// 编辑商品类型页面
		adminRouters.GET("/goodsType/edit", admin.GoodsTypeController{}.Edit)
		// 编辑商品类型
		adminRouters.PUT("/goodsType/edit", admin.GoodsTypeController{}.DoEdit)
		// 删除商品类型页面
		adminRouters.DELETE("/goodsType/delete", admin.GoodsTypeController{}.Delete)

		// 商品类型属性页面
		adminRouters.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController{}.Index)
		// 添加商品类型属性页面
		adminRouters.GET("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.Add)
		// 添加商品类型属性
		adminRouters.POST("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.DoAdd)
		// 编辑商品类型属性页面
		adminRouters.GET("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeController{}.Edit)
		// 编辑商品类型属性
		adminRouters.PUT("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeController{}.DoEdit)
		// 删除商品类型属性页面
		adminRouters.DELETE("/goodsTypeAttribute/delete", admin.GoodsTypeAttributeController{}.Delete)

		// 商品颜色页面
		adminRouters.GET("/goodsColor", admin.GoodsColorController{}.Index)
		// 添加商品颜色页面
		adminRouters.GET("/goodsColor/add", admin.GoodsColorController{}.Add)
		// 添加商品颜色
		adminRouters.POST("/goodsColor/add", admin.GoodsColorController{}.DoAdd)
		// 编辑商品颜色页面
		adminRouters.GET("/goodsColor/edit", admin.GoodsColorController{}.Edit)
		// 编辑商品颜色
		adminRouters.PUT("/goodsColor/edit", admin.GoodsColorController{}.DoEdit)
		// 删除商品颜色页面
		adminRouters.DELETE("/goodsColor/delete", admin.GoodsColorController{}.Delete)

		// 商品页面
		adminRouters.GET("/goods", admin.GoodsController{}.Index)
		// 添加商品页面
		adminRouters.GET("/goods/add", admin.GoodsController{}.Add)
		// 添加商品
		adminRouters.POST("/goods/add", admin.GoodsController{}.DoAdd)
		// 编辑商品页面
		adminRouters.GET("/goods/edit", admin.GoodsController{}.Edit)
		// 编辑商品
		adminRouters.PUT("/goods/edit", admin.GoodsController{}.DoEdit)
		// wysiwyg-editor上传图片
		adminRouters.POST("/goods/imageUpload", admin.GoodsController{}.ImageUpload)
		// Ajax动态获取商品类型属性
		adminRouters.GET("/goods/goodsTypeAttribute", admin.GoodsController{}.GoodsTypeAttribute)

	}
}
