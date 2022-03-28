package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con FocusController) Add(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con FocusController) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con FocusController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
