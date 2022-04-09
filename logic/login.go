package logic

import (
	"ziweiShop/dao/mysql"
	"ziweiShop/models"
	"ziweiShop/pkg/captcha"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
)

type LoginLogic struct {
}

func (LoginLogic) DoLogin(c *gin.Context, p *models.LoginParams) (err error) {
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

	// 根据username查询userId
	ManagerInfo := mysql.GetManagerInfoByUsername(p.Username)
	if ManagerInfo == nil {
		return ErrorManagerNotExist
	}
	// 设置sessions、cookie
	session := sessions.Default(c)
	session.Set("username", p.Username)
	session.Set("managerId", ManagerInfo.Id)
	session.Set("roleId", ManagerInfo.RoleId)
	err = session.Save()
	return
}

func (LoginLogic) Logout(c *gin.Context) (err error) {
	// 删除sessions
	session := sessions.Default(c)
	session.Delete("username")
	return session.Save()
}

func (LoginLogic) GenCaptcha() (map[string]string, error) {
	id, b64s, err := captcha.GenCaptcha()
	data := map[string]string{
		"captcha_id":  id,
		"captcha_url": b64s,
	}
	return data, err
}
