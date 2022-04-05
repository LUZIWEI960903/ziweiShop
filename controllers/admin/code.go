package admin

import "errors"

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
	CodeEmptyTitle
	CodeEmptyDecription
	CodeRoleExist
	CodeAddRoleErr
	CodeGetRoleListErr
	CodeGetRoleErr
	CodeEditRoleErr
	CodeDeleteRoleErr
	CodeManagerExistErr
	CodeRoleNotExistErr
	CodeAddManagerErr
	CodeGetIndexManagerListErr
	CodeGetEditManagerErr
	CodeManagerDoEditErr
	CodeDeleteManagerErr
	CodeGetTopAccessListErr
)

var (
	ErrEmptyTitle = errors.New("Title can not empty")
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:                "",
	CodeServerBusy:             "Server busy",
	CodeInValidParams:          "Invalid params",
	CodeManagerNotExist:        "Manager is not exist",
	CodeManagerPasswordErr:     "Manager username or password error",
	CodeInValidCaptcha:         "Captcha verify error",
	CodeLogoutError:            "Logout error",
	CodeGenCaptchaError:        "Gen captcha error",
	CodeNeedToLogin:            "Please login in",
	CodeEmptyTitle:             "Title can not empty",
	CodeEmptyDecription:        "Decription can not empty",
	CodeRoleExist:              "Role exist",
	CodeAddRoleErr:             "Add role error",
	CodeGetRoleListErr:         "Get roleList error",
	CodeGetRoleErr:             "Get role info error",
	CodeEditRoleErr:            "Edit role error",
	CodeDeleteRoleErr:          "Delete role error",
	CodeManagerExistErr:        "Manager is exist",
	CodeRoleNotExistErr:        "Role not exist",
	CodeAddManagerErr:          "Add manager error",
	CodeGetIndexManagerListErr: "Get index manager list error",
	CodeGetEditManagerErr:      "Get edit manager info error",
	CodeManagerDoEditErr:       "Manager do edit error",
	CodeDeleteManagerErr:       "Delete manager error",
	CodeGetTopAccessListErr:    "Get top access list error",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
