package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainPageController struct {
}

func (con MainPageController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con MainPageController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
