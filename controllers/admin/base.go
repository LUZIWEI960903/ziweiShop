package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) success(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func (con BaseController) error(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
