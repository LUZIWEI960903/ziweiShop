package logic

import "errors"

var (
	ErrorManagerExist    = errors.New("Manager is exist~~")
	ErrorManagerNotExist = errors.New("Manager is not exist~~")
	ErrorManagerPassword = errors.New("Manager username or password failed~~")
	ErrorInValidCaptcha  = errors.New("Invaild captcha~~")

	ErrorRoleExist    = errors.New("Role is exist~~")
	ErrorRoleNotExist = errors.New("Role is not exist~~")
	ErrorAddRole      = errors.New("Add Role failed~~")
	ErrorGetRole      = errors.New("Get Role info failed~~")
	ErrorEditRole     = errors.New("Edit Role failed~~")
)
