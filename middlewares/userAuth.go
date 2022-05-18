package middlewares

import (
	"net/http"
	"regexp"
	"ziweiShop/controllers/admin"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

	"github.com/gin-gonic/gin"
)

func InitUserAuthMiddleware(c *gin.Context) {
	userInfo := models.User{}
	cookie.Cookie.Get(c, "userInfo", &userInfo)
	pattern := `.*?@.*?`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(userInfo.Email) {
		c.JSON(http.StatusOK, gin.H{
			"errcode": admin.CodeGetAccessErr,
			"errmsg":  admin.CodeGetAccessErr.Msg(),
		})
		return
	}
}
