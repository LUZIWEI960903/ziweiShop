package admin

type ResCode int

const (
	CodeSuccess ResCode = iota
	CodeServerBusy

	CodeInValidParams = 1000 + iota
	CodeManagerNotExist
	CodeManagerPasswordErr
	CodeInValidCaptcha
	CodeLogoutError
	CodeGenCaptchaError
	CodeNeedToLogin
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:            "",
	CodeServerBusy:         "Server busy",
	CodeInValidParams:      "Invalid params",
	CodeManagerNotExist:    "Manager is not exist",
	CodeManagerPasswordErr: "Manager username or password error",
	CodeInValidCaptcha:     "Captcha verify failed",
	CodeLogoutError:        "Logout error",
	CodeGenCaptchaError:    "Gen captcha failed",
	CodeNeedToLogin:        "Please login in",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
