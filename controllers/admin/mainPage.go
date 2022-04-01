package admin

import (
	"github.com/gin-gonic/gin"
)

type MainPageController struct {
	BaseController
}

func (con MainPageController) Index(c *gin.Context) {
	con.success(c, true)
}
