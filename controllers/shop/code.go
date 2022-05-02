package shop

type ResCode int

const (
	CodeSuccess ResCode = iota
	CodeServerBusy
	CodeInvalidAccess

	CodeInValidParams ResCode = 1000 + iota

	CodeGetAddPageDataErr
	CodeDoAddLogicErr
	CodeGetIndexPageDataErr
	CodeGetEditPageDataErr
	CodeDoEditLogicErr
	CodeDeleteLogicErr
	CodeAjaxErr
)

var CodeMsgMap = map[ResCode]string{
	CodeSuccess:       "",
	CodeServerBusy:    "Server busy",
	CodeInvalidAccess: "Invalid access",
	CodeInValidParams: "Invalid params",

	CodeGetAddPageDataErr:   "Get add page data error",
	CodeDoAddLogicErr:       "Do add logic error",
	CodeGetIndexPageDataErr: "Get index page data error",
	CodeGetEditPageDataErr:  "Get edit page data error",
	CodeDoEditLogicErr:      "Do edit logic error",
	CodeDeleteLogicErr:      "Delete logic error",
	CodeAjaxErr:             "Ajax error",
}

func (code ResCode) Msg() string {
	msg, ok := CodeMsgMap[code]
	if !ok {
		msg = CodeMsgMap[CodeServerBusy]
	}
	return msg
}
