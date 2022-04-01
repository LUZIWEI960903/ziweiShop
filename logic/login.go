package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/captcha"
)

type LoginLogic struct {
}

func (LoginLogic) DoLogin(p *models.LoginParams) (err error) {
	// 校验管理员是否存在
	err = mysql.IsManagerExist(p.Username)
	if err != nil {
		return ErrorManagerNotExist
	}
	// 校验管理员账号+密码
	err = mysql.VerifyUsernameAndPassword(p.Username, p.Password)
	if err != nil {
		return ErrorManagerPassword
	}
	// 校验验证码
	if ok := captcha.VerifyCaptcha(p.CaptchaId, p.CaptchaValue); !ok {
		return ErrorInValidCaptcha
	}
	return
}
