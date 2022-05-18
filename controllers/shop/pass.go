package shop

import (
	"fmt"
	"ziweiShop/logic"
	"ziweiShop/models"
	"ziweiShop/pkg/cookie"

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

// Login 用户登录页面的接口
func (con PassController) Login(c *gin.Context) {
	con.success(c, true)
}

// DoLogin 用户登录的接口
func (con PassController) DoLogin(c *gin.Context) {
	// 解析参数
	p := new(models.UserLoginParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: shop] [func: (con PassController) DoLogin(c *gin.Context)] [c.ShouldBindJSON(p)] failed, ", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}

	// 获取上一个url地址
	prePage := c.Request.Referer()

	// 业务逻辑
	ok := logic.PassLogic{}.DoLogin(c, p)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) DoLogin(c *gin.Context)] [logic.PassLogic{}.DoLogin(c, p)] failed")
		con.error(c, CodeDoLoginErr)
		return
	}
	con.success(c, prePage)
}

// Logout 用户退出登录的接口
func (con PassController) Logout(c *gin.Context) {
	ok := logic.PassLogic{}.Logout(c)
	if !ok {
		zap.L().Error("[pkg: shop] [func: (con PassController) Logout(c *gin.Context)] [logic.PassLogic{}.Logout(c)] failed")
		con.error(c, CodeDoLogoutErr)
		return
	}
	con.success(c, true)
}

// Checkout 确认订单信息页面的接口
func (con PassController) Checkout(c *gin.Context) {
	user := models.User{}
	cookie.Cookie.Get(c, "userInfo", &user)
	fmt.Printf("%#v\n", user)
}
