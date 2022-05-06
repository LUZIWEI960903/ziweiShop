package shop

import (
	"net/http"
	"strconv"

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

func (BaseController) getPagePageSize(c *gin.Context) (page, pagesize int, err error) {
	page, err1 := strconv.Atoi(c.DefaultQuery("p", "1"))
	if err1 != nil {
		return 0, 0, err1
	}
	pageSize, err2 := strconv.Atoi(c.DefaultQuery("s", "10"))
	if err2 != nil {
		return 0, 0, err2
	}
	return page, pageSize, nil
}
