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
	data, ok := logic.PassLogic{}.DoRegister1(c, p)
	if !ok || data == "" {
		zap.L().Error("[pkg: shop] [func: (con PassController) Register1(c *gin.Context)] [logic.PassLogic{}.Register1(c,p)] failed")
		con.error(c, CodeRegisterErr)
		return
	}
	con.success(c, data)
}

// Register2 用户注册第二步页面的接口
func (con PassController) Register2(c *gin.Context) {
	// 解析参数
	sign := c.Query("sign")
	verifyCode := c.Query("verifyCode") // 注册第一步输入的captchaValue

	// 业务逻辑
	ok := logic.PassLogic{}.Register2(c, sign, verifyCode)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) Register2(c *gin.Context)] [logic.PassLogic{}.Register2(c, sign, verifyCode)] failed")
		con.error(c, CodeInValidParams)
		return
	}
	con.success(c, true)
}

// DoRegister2 用户注册第二步的接口
func (con PassController) DoRegister2(c *gin.Context) {
	// 解析参数
	sign := c.Query("sign")
	msgCode := c.Query("msgCode")

	// 业务逻辑
	ok := logic.PassLogic{}.DoRegister2(c, sign, msgCode)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) DoRegister2(c *gin.Context)] [logic.PassLogic{}.DoRegister2(c, sign, msgCode)] failed")
		con.error(c, CodeInValidParams)
		return
	}
	con.success(c, true)
}

// Register3 用户注册第三步页面的接口
func (con PassController) Register3(c *gin.Context) {
	// 解析参数
	sign := c.Query("sign")
	msgCode := c.Query("msgCode")

	// 业务逻辑
	ok := logic.PassLogic{}.Register3(c, sign, msgCode)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) Register3(c *gin.Context)] [logic.PassLogic{}.Register3(c, sign, msgCode)] failed")
		con.error(c, CodeInValidParams)
		return
	}

	con.success(c, true)
}

// DoRegister3 用户注册第三步的接口
func (con PassController) DoRegister3(c *gin.Context) {
	// 解析参数
	p := new(models.Register3Params)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: shop] [func: (con PassController) DoRegister3(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 业务逻辑
	ok := logic.PassLogic{}.DoRegister3(c, p)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) DoRegister3(c *gin.Context)] [logic.PassLogic{}.DoRegister3(c,p)] failed")
		con.error(c, CodeRegisterErr)
		return
	}

	con.success(c, true)
}
