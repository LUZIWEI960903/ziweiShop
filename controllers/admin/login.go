package admin

import (
	"errors"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Login(c *gin.Context) {
	con.success(c, NeedToLogin)
}

// DoLogin 管理员登录的接口
func (con LoginController) DoLogin(c *gin.Context) {
	// 获取登录表单信息
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, InValidParams)
		return
	}
	// 登录业务逻辑
	err := logic.LoginLogic{}.DoLogin(c, p)
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

// Logout 管理员登出的接口
func (con LoginController) Logout(c *gin.Context) {
	err := logic.LoginLogic{}.Logout(c)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con LoginController) Logout(c *gin.Context)] [logic.LoginLogic{}.Logout(c)] failed, err:", zap.Error(err))
		con.error(c, LogoutError)
		return
	}
	con.success(c, LogoutSuccess)
}

// Captcha 获取验证码的接口
func (con LoginController) Captcha(c *gin.Context) {
	var captchaServie = logic.LoginLogic{}
	data, err := captchaServie.GenCaptcha()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) Captcha(c *gin.Context)] [captcha.GenCaptcha()] failed, err:", zap.Error(err))
		con.error(c, GenCaptchaError)
		return
	}
	con.success(c, GenCaptchaSuccess, data)
}
