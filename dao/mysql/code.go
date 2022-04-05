package mysql

import "errors"

var (
	ErrManagerExist    = errors.New("Manager is exist~~")
	ErrManagerNotExist = errors.New("Manager is not exist~~")
	ErrManagerPassword = errors.New("Manager username or password failed~~")
	ErrGetManagerList  = errors.New("Get managerList failed~~")

	ErrRoleExist = errors.New("Role is not exist~~")
	ErrNoRole    = errors.New("No role exist~~")

	ErrGetTopAccessList = errors.New("Get top access list error~~")
)
