package logic

import "errors"

var (
	ErrorManagerNotExist = errors.New("Manager is not exist~~")
	ErrorManagerPassword = errors.New("Manager username or password failed~~")
	ErrorInValidCaptcha  = errors.New("Invaild captcha~~")
)
