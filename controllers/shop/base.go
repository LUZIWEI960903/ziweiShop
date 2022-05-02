package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (con BaseController) success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": CodeSuccess,
		"errmsg":  CodeSuccess.Msg(),
		"data":    data,
	})
}

func (con BaseController) error(c *gin.Context, code ResCode) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": code,
		"errmsg":  code.Msg(),
	})
}
