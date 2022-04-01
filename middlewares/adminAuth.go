package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	// 获取url
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// 获取session
	session := sessions.Default(c)
	username, ok := session.Get("username").(string)
	if !ok || username == "" {
		if pathname != "/admin/login" && pathname != "/admin/captcha" {
			c.Redirect(http.StatusFound, "/admin/login")
		}
	}
	fmt.Printf("Welcome %s~~\n", username)
}
