package mysql

import "errors"

var (
	ErrManagerExist    = errors.New("Manager is exist~~")
	ErrManagerNotExist = errors.New("Manager is not exist~~")
	ErrManagerPassword = errors.New("Manager username or password failed~~")
	ErrGetManagerList  = errors.New("Get managerList failed~~")

	ErrRoleExist = errors.New("Role is not exist~~")
	ErrNoRole    = errors.New("No role exist~~")

	ErrGetTopAccessList               = errors.New("Get top access list error~~")
	ErrGetTopAccessListWithAccessList = errors.New("Get top access list with access list error~~")
	ErrGetAccess                      = errors.New("Get access info error~~")

	ErrGetFocus = errors.New("Get focus error~~")

	ErrAjaxChangeStatus = errors.New("Update status error~~")

	ErrGetGoodsCate = errors.New("Get goods cate info error~~")

	ErrGetGoodsType = errors.New("Get goods type list error~~")

	ErrGetGoodsTypeAttribute = errors.New("Get goods type attribute error~~")

	ErrGetData = errors.New("Get data error~~")

	ErrGetGoodsColor = errors.New("Get goods color error~~")

	ErrGetGoods = errors.New("Get goods error~~")

	ErrGetGoodsImage = errors.New("Get goods image error~~")

	ErrGetNav = errors.New("Get nav error~~")

	ErrEmailExist = errors.New("Email exist~~")
)
