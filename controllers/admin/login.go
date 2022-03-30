package admin

import (
	"fmt"
	"net/http"
	"ziweiShop/models"
	"ziweiShop/pkg/captcha"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.JSON(http.StatusOK, nil)
}

func (con LoginController) DoLogin(c *gin.Context) {
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		c.String(http.StatusOK, "Login failed~~")
		return
	}
	if ok := captcha.VerifyCaptcha(p.CaptchaId, p.CaptchaValue); !ok {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [captcha.VerifyCaptcha(p.CaptchaId,p.CaptchaValue)] failed")
		c.String(http.StatusOK, "Captcha failed~~")
		return
	}
	fmt.Println(p)
	c.JSON(http.StatusOK, "Login success~~")
}

// 获取验证码接口
func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := captcha.GenCaptcha()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) Captcha(c *gin.Context)] [captcha.GenCaptcha()] failed, err:", zap.Error(err))
		c.String(http.StatusOK, "Gen captcha failed~~")
	}
	c.JSON(http.StatusOK, gin.H{
		"id":   id,
		"b64s": b64s,
	})
}
