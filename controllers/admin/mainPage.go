package admin

import (
	"github.com/gin-gonic/gin"
)

type MainPageController struct {
	BaseController
}

// Index 显示后台主页面的接口
func (con MainPageController) Index(c *gin.Context) {
	con.success(c, true)
}
