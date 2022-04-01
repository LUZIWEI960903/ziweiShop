package mysql

import "errors"

var (
	ErrManagerExist    = errors.New("Manager is not exist~~")
	ErrManagerPassword = errors.New("Manager username or password failed~~")

	ErrRoleExist = errors.New("Role is not exist~~")
	ErrNoRole    = errors.New("No role exist~~")
)
