package logic

import (
	"regexp"
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/captcha"
	"ziweiShop/pkg/email"
	"ziweiShop/pkg/tools"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type PassLogic struct {
}

func (PassLogic) DoRegister1(c *gin.Context, p *models.Register1Params) (string, bool) {
	// 校验验证码
	if len(p.CaptchaValue) != 4 || !captcha.VerifyCaptcha(p.CaptchaId, p.CaptchaValue) {
		return "", false
	}

	// 将验证码captchaValue保存至session
	captchaSession := sessions.Default(c)
	captchaSession.Set("captchaValue", p.CaptchaValue)
	captchaSession.Save()

	// 校验邮箱是否合法
	pattern := `.*?@.*?`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(p.Email) {
		return "", false
	}

	// 查询该邮箱是否已存在
	if mysql.IsEmailExist(p.Email) {
		return "", false
	}

	// 判断当前ip地址今天注册邮箱的个数，同个ip同一天最多注册10个
	currentDay := tools.GetDate()
	ip := c.ClientIP()
	if !mysql.IsValidRegisterCount(ip, currentDay) {
		return "", false
	}
	sign := tools.MD5(ip + currentDay)
	// 查询当前邮箱在今天发送了几次验证码
	sendMsgCount := mysql.GetSendMsgCount(p.Email, currentDay)
	if sendMsgCount == 0 { // 代表该邮箱今天还没发送过验证码
		msgCode := tools.GenerateCode()                           // 生成验证码
		if err := email.SendEmail(p.Email, msgCode); err != nil { // 发送验证码
			return "", false
		}
		session := sessions.Default(c) // 将验证码保存至session
		session.Set("msgCode", msgCode)
		session.Save()
		// 创建sign
		if err := mysql.CreateEmailTemp(p.Email, currentDay, ip, sign); err != nil { // 创建一条数据记录
			return "", false
		}
	} else { // 代表该邮箱今天已经发送过验证码了，需要判断发送次数是否小于3次
		userTempInfo, err := mysql.GetUserTempInfo(p.Email, currentDay)
		if err != nil {
			return "", false
		}
		if userTempInfo.SendCount < 3 { // 发送验证码+更新次数
			msgCode := tools.GenerateCode()                           // 生成验证码
			if err := email.SendEmail(p.Email, msgCode); err != nil { // 发送验证码
				return "", false
			}
			session := sessions.Default(c) // 将验证码保存至session
			session.Set("msgCode", msgCode)
			session.Save()
			if err := mysql.UpdateMsgCount(&userTempInfo); err != nil {
				return "", false
			}
		} else {
			return "", false
		}
	}
	return sign, true
}

func (PassLogic) Register2(c *gin.Context, sign, verifyCode string) bool {
	// 校验verifyCode
	session := sessions.Default(c)
	sessionVerifyCode, ok1 := session.Get("captchaValue").(string)
	if !ok1 || sessionVerifyCode != verifyCode {
		return false
	}

	// 校验sign
	_, ok := mysql.GetUserTempBySign(sign)
	if !ok {
		return false
	}
	return true
}

func (PassLogic) DoRegister2(c *gin.Context, sign, smsCode string) bool {
	// 校验smsCode
	session := sessions.Default(c)
	sessionSmsCode, ok1 := session.Get("msgCode").(string)
	if !ok1 || sessionSmsCode != smsCode {
		return false
	}

	// 校验sign
	userTemp, ok := mysql.GetUserTempBySign(sign)
	if !ok {
		return false
	}

	// 校验验证码是否过期
	nowTime := tools.GetUnix()
	if (nowTime - int64(userTemp.AddTime)) > 15*60 {
		return false
	}
	return true
}
