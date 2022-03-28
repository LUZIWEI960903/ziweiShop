package routers

import "github.com/gin-gonic/gin"

func adminRoutersInit(r *gin.Engine) {
	//adminRouters := r.Group("/admin", func(c *gin.Context) {
	//	c.String(200, "666")
	//})
	r.GET("/", func(c *gin.Context) {
		c.String(200, "666")
	})
}
