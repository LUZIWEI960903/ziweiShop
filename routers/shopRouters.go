package routers

import (
	"ziweiShop/controllers/shop"
	"ziweiShop/middlewares"

	"github.com/gin-gonic/gin"
)

func shopRoutersInit(r *gin.Engine) {
	shopRouters := r.Group("/")
	{
		// 商城主页面
		shopRouters.GET("/", shop.IndexController{}.Index)
		// 商品分类页面
		shopRouters.GET("/category:id", shop.ProductController{}.Category)
		// 商品详情页面
		shopRouters.GET("/detail", shop.ProductController{}.Detail)
		// Ajax点击颜色获取对应商品的图片
		shopRouters.GET("/product/getImgList", shop.ProductController{}.GetImgList)

		// 获取购物车数据
		shopRouters.GET("/cart", shop.CartController{}.Get)
		// 添加购物车
		shopRouters.GET("/cart/addCart", shop.CartController{}.AddCart)
		// 添加购物车成功
		shopRouters.GET("/cart/addCartSuccess", shop.CartController{}.AddCartSuccess)
		// 增加购物车商品数量
		shopRouters.GET("/cart/incCart", shop.CartController{}.IncCart)
		// 减少购物车商品数量
		shopRouters.GET("/cart/decCart", shop.CartController{}.DecCart)
		// 改变购物车指定商品选择的状态
		shopRouters.GET("/cart/changeOneCart", shop.CartController{}.ChangeOneCart)
		// 改变购物车所有商品选择的状态
		shopRouters.GET("/cart/changeAllCart", shop.CartController{}.ChangeAllCart)
		// 删除购物车指定商品
		shopRouters.DELETE("/cart/deleteCart", shop.CartController{}.DeleteCart)

		// 生成验证码
		shopRouters.GET("pass/captcha", shop.PassController{}.Captcha)
		// 用户注册页面1
		shopRouters.GET("/pass/register1", shop.PassController{}.Register1)
		// 用户注册页面1过程
		shopRouters.POST("/pass/register1", shop.PassController{}.DoRegister1)
		// 用户注册页面2
		shopRouters.GET("/pass/register2", shop.PassController{}.Register2)
		// 用户注册页面2过程
		shopRouters.POST("/pass/register2", shop.PassController{}.DoRegister2)
		// 用户注册页面3
		shopRouters.GET("/pass/register3", shop.PassController{}.Register3)
		// 用户注册页面3过程
		shopRouters.POST("/pass/register3", shop.PassController{}.DoRegister3)

		// 登录界面
		shopRouters.GET("/pass/login", shop.PassController{}.Login)
		// 执行登录
		shopRouters.POST("/pass/login", shop.PassController{}.DoLogin)
		// 执行登出
		shopRouters.GET("/pass/logout", shop.PassController{}.Logout)

		// 确认订单信息
		shopRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, shop.BuyController{}.Checkout)

		// 增加收货地址
		shopRouters.POST("/address/addAddress", middlewares.InitUserAuthMiddleware, shop.AddressController{}.AddAddress)

	}
}
