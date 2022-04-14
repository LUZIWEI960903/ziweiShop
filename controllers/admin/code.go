package admin

import "errors"

type ResCode int

const (
	CodeSuccess ResCode = iota
	CodeServerBusy
	CodeInvalidAccess

	CodeInValidParams ResCode = 1000 + iota
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
	CodeEmptyModuleName
	CodeAddAccessErr
	CodeTopAccessListWithAccessListErr
	CodeGetAccessErr
	CodeEditAccessErr
	CodeDeleteAccessErr
	CodeRoleAuthAccessErr
	CodeGetAdminIndexPageInfoErr
	CodeGetFocusErr
	CodeUploadImgErr
	CodeAddFocusErr
	CodeGetFocusListErr
	CodeEditFocusErr
	CodeDeleteFocusErr
	CodeAjaxChangeStatusErr
	CodeAjaxChangeSortErr
	CodeAddGoodsCateErr
	CodeGetTopGoodsCateListErr
	CodeGetTopGoodsCateWithGoodsCateListErr
	CodeGetGoodsCateInfoErr
	CodeEditGoodsCateErr
	CodeDeleteGoodsCateErr
	CodeAddGoodsTypeErr
	CodeGetGoodsTypeListErr
	CodeGetGoodsTypeInfoErr
	CodeEditGoodsTypeErr
)

var (
	ErrEmptyTitle      = errors.New("Title can not empty")
	ErrEmptyModuleName = errors.New("Empty module name")
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:                             "",
	CodeServerBusy:                          "Server busy",
	CodeInvalidAccess:                       "Invalid access",
	CodeInValidParams:                       "Invalid params",
	CodeManagerNotExist:                     "Manager is not exist",
	CodeManagerPasswordErr:                  "Manager username or password error",
	CodeInValidCaptcha:                      "Captcha verify error",
	CodeLogoutError:                         "Logout error",
	CodeGenCaptchaError:                     "Gen captcha error",
	CodeNeedToLogin:                         "Please login in",
	CodeEmptyTitle:                          "Title can not empty",
	CodeEmptyDecription:                     "Decription can not empty",
	CodeRoleExist:                           "Role exist",
	CodeAddRoleErr:                          "Add role error",
	CodeGetRoleListErr:                      "Get roleList error",
	CodeGetRoleErr:                          "Get role info error",
	CodeEditRoleErr:                         "Edit role error",
	CodeDeleteRoleErr:                       "Delete role error",
	CodeManagerExistErr:                     "Manager is exist",
	CodeRoleNotExistErr:                     "Role not exist",
	CodeAddManagerErr:                       "Add manager error",
	CodeGetIndexManagerListErr:              "Get index manager list error",
	CodeGetEditManagerErr:                   "Get edit manager info error",
	CodeManagerDoEditErr:                    "Manager do edit error",
	CodeDeleteManagerErr:                    "Delete manager error",
	CodeGetTopAccessListErr:                 "Get top access list error",
	CodeEmptyModuleName:                     "Empty module name",
	CodeAddAccessErr:                        "Add access error",
	CodeTopAccessListWithAccessListErr:      "Get top access list with access list error",
	CodeGetAccessErr:                        "Get access error",
	CodeEditAccessErr:                       "Edit access error",
	CodeDeleteAccessErr:                     "Delete access error",
	CodeRoleAuthAccessErr:                   "Role auth access error",
	CodeGetAdminIndexPageInfoErr:            "Get admin index page error",
	CodeGetFocusErr:                         "Get focus error",
	CodeUploadImgErr:                        "Upload image error",
	CodeAddFocusErr:                         "Add focus error",
	CodeGetFocusListErr:                     "Get focus list error",
	CodeEditFocusErr:                        "Edit focus error",
	CodeDeleteFocusErr:                      "Delete focus error",
	CodeAjaxChangeStatusErr:                 "Ajax change status error",
	CodeAjaxChangeSortErr:                   "Ajax change sort error",
	CodeAddGoodsCateErr:                     "Add goods cate error",
	CodeGetTopGoodsCateListErr:              "Get top goods cate list error",
	CodeGetTopGoodsCateWithGoodsCateListErr: "Get top goods cate with goods cate list error",
	CodeGetGoodsCateInfoErr:                 "Get goods cate info error",
	CodeEditGoodsCateErr:                    "Edit goods cate error",
	CodeDeleteGoodsCateErr:                  "Delete goods cate error",
	CodeAddGoodsTypeErr:                     "Add goods type error",
	CodeGetGoodsTypeListErr:                 "Get goods type list error",
	CodeGetGoodsTypeInfoErr:                 "Get goods type info error",
	CodeEditGoodsTypeErr:                    "Edit goods type error",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
