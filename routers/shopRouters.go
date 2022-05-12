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

	}
}
