package email

import (
	"net/smtp"
)

const (
	SENDEMAIL = "626398871@qq.com"
	EMAILKEY  = "nehiwboppcxqbcje"
)

func SendEmail(receiverEmail, msgCode string) error {
	// 设置发送方账户信息信息，这边密码要填入邮箱授权码，获取授权码步骤可以网上查一查
	auth := smtp.PlainAuth("", SENDEMAIL, EMAILKEY, "smtp.qq.com")
	// 设置接收方信息
	to := []string{receiverEmail}
	// 编辑邮件内容
	msg := []byte("To: " + receiverEmail + "\r\n" +
		"Subject: ziweiShop register!\r\n" +
		"\r\n" +
		"Register success!!\r\nYou code:" + msgCode)
	// 调用函数发送邮件
	return smtp.SendMail("smtp.qq.com:25", auth, SENDEMAIL, to, msg)
}
