package shop

import (
	"ziweiShop/logic"
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PassController struct {
	BaseController
}

// Captcha 获取验证码的接口
func (con PassController) Captcha(c *gin.Context) {
	data, err := logic.LoginLogic{}.GenCaptcha()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con PassController) Captcha(c *gin.Context)] [logic.LoginLogic{}.GenCaptcha()] failed, err:", zap.Error(err))
		con.error(c, CodeGenCaptchaError)
		return
	}
	con.success(c, data)
}

// Register1 用户注册第一步页面的接口
func (con PassController) Register1(c *gin.Context) {
	con.success(c, true)
}

// DoRegister1 用户注册第一步的接口
func (con PassController) DoRegister1(c *gin.Context) {
	// 解析参数
	p := new(models.Register1Params)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: shop] [func: (con PassController) Register1(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	ok := logic.PassLogic{}.Register1(p)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) Register1(c *gin.Context)] [logic.PassLogic{}.Register1(p)] failed")
		con.error(c, CodeRegisterErr)
		return
	}
	con.success(c, true)
}
