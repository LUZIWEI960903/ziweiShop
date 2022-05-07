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

	}
}
