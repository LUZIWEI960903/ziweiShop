package admin

import (
	"errors"
	logic2 "ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

// Index 管理员登录页面的接口
func (con LoginController) Index(c *gin.Context) {
	con.success(c, true)
}

// DoLogin 管理员登录的接口
func (con LoginController) DoLogin(c *gin.Context) {
	// 获取登录表单信息
	p := new(models.LoginParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 登录业务逻辑
	err := logic2.LoginLogic{}.DoLogin(c, p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) DoLogin(c *gin.Context)] [logiclogic.LoginService{}.DoLogin(p)] failed, err:", zap.Error(err))
		// 用户不存在
		if errors.Is(err, logic2.ErrorManagerNotExist) {
			con.error(c, CodeManagerNotExist)
			return
		}
		// 用户+密码错误
		if errors.Is(err, logic2.ErrorManagerPassword) {
			con.error(c, CodeManagerPasswordErr)
			return
		}
		// 验证码错误
		if errors.Is(err, logic2.ErrorInValidCaptcha) {
			con.error(c, CodeInValidCaptcha)
			return
		}
		// 其余错误
		con.error(c, CodeServerBusy)
		return
	}
	// 登录成功
	con.success(c, true)
}

// Logout 管理员登出的接口
func (con LoginController) Logout(c *gin.Context) {
	err := logic2.LoginLogic{}.Logout(c)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con LoginController) Logout(c *gin.Context)] [logic.LoginLogic{}.Logout(c)] failed, err:", zap.Error(err))
		con.error(c, CodeLogoutError)
		return
	}
	con.success(c, true)
}

// Captcha 获取验证码的接口
func (con LoginController) Captcha(c *gin.Context) {
	var captchaServie = logic2.LoginLogic{}
	data, err := captchaServie.GenCaptcha()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con LoginController) Captcha(c *gin.Context)] [captcha.GenCaptcha()] failed, err:", zap.Error(err))
		con.error(c, CodeGenCaptchaError)
		return
	}
	con.success(c, data)
}
