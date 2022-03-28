package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "admin do login")
}
