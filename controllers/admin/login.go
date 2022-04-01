package admin

import (
	"errors"
	"net/http"
	"ziweiShop/logic"
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
	// 获取登录表单信息
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, InValidParams)
		return
	}
	// 登录业务逻辑
	err := logic.LoginLogic{}.DoLogin(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [logicAdmin.LoginService{}.DoLogin(p)] failed, err:", zap.Error(err))
		// 用户不存在
		if errors.Is(err, logic.ErrorManagerNotExist) {
			con.error(c, ManagerNotExist)
			return
		}
		// 用户+密码错误
		if errors.Is(err, logic.ErrorManagerPassword) {
			con.error(c, ManagerPassword)
			return
		}
		// 验证码错误
		if errors.Is(err, logic.ErrorInValidCaptcha) {
			con.error(c, InValidCaptcha)
			return
		}
		// 其余错误
		con.error(c, ServiceBusy)
		return
	}
	// 登录成功
	con.success(c, LoginSuccess)
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
