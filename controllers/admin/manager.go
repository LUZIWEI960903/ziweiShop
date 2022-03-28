package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con ManagerController) Add(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con ManagerController) Edit(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con ManagerController) Delete(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}
