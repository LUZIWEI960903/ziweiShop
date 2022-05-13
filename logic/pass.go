package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/captcha"
)

type PassLogic struct {
}

func (PassLogic) Register1(p *models.Register1Params) bool {
	// 查询该邮箱是否已存在
	if mysql.IsEmailExist(p.Email) {
		return false
	}

	// 校验验证
	if !captcha.VerifyCaptcha(p.CaptchaId, p.CaptchaValue) {
		return false
	}
	return true
}
