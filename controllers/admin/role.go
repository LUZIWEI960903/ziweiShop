package admin

import (
	"errors"
	"fmt"
	"strconv"
	logic2 "ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

// Index 显示管理员角色页面的接口
func (con RoleController) Index(c *gin.Context) {
	var roleService = logic2.RoleLogic{}
	roleList, err := roleService.GetRoleList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Index(c *gin.Context)] [roleService.GetRoleList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetRoleListErr)
		return
	}
	con.success(c, roleList)
}

// Add 给管理员增加角色页面的接口
func (con RoleController) Add(c *gin.Context) {
	con.success(c, true)
}

// DoAdd 给管理员增加角色的接口
func (con RoleController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddRoleParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 参数校验
	if err := verifyAddRoleParams(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAdd(c *gin.Context)] [verifyAddRoleParams(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEmptyTitle)
		return
	}
	// 业务逻辑
	err := logic2.RoleLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAdd(c *gin.Context)] [logic.RoleLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		if errors.Is(err, logic2.ErrorRoleExist) {
			con.error(c, CodeRoleExist)
			return
		}
		if errors.Is(err, logic2.ErrorAddRole) {
			con.error(c, CodeAddRoleErr)
			return
		}
	}
	con.success(c, true)
}

// Edit 给管理员编辑角色页面的接口
func (con RoleController) Edit(c *gin.Context) {
	// 解析参数
	roleIdStr := c.Query("id")
	// 校验参数
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Edit(c *gin.Context)] [strconv.Atoi(roleIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	var roleService = logic2.RoleLogic{}
	data, err := roleService.GetRoleById(roleId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Edit(c *gin.Context)] [roleService.GetRoleById(roleId)] failed, err:", zap.Error(err))
		con.error(c, CodeGetRoleErr)
		return
	}
	con.success(c, data)
}

// DoEdit 给管理员编辑角色的接口
func (con RoleController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditRoleParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 参数校验
	if err := verifyEditRoleParams(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoEdit(c *gin.Context)] [verifyAddRoleParams(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEmptyTitle)
		return
	}
	// 业务逻辑
	err := logic2.RoleLogic{}.DoEdit(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoEdit(c *gin.Context)] [logic.RoleLogic{}.DoEdit(p)] failed, err:", zap.Error(err))
		if errors.Is(err, logic2.ErrorRoleExist) {
			con.error(c, CodeRoleExist)
			return
		}
		if errors.Is(err, logic2.ErrorEditRole) {
			con.error(c, CodeEditRoleErr)
			return
		}
	}
	con.success(c, true)
}

// Delete 给管理员删除角色的接口
func (con RoleController) Delete(c *gin.Context) {
	// 解析参数
	roleIdStr := c.Query("id")
	// 校验参数
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Delete(c *gin.Context)] [strconv.Atoi(roleIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err = logic2.RoleLogic{}.DeleteRoleById(roleId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Delete(c *gin.Context)] [logic.RoleLogic{}.DeleteRoleById(roleId)] failed, err:", zap.Error(err))
		con.error(c, CodeDeleteRoleErr)
		return
	}
	con.success(c, true)
}

// Auth 给角色授权页面的接口
func (con RoleController) Auth(c *gin.Context) {
	// 解析参数
	roleIdStr := c.Query("id")
	roleId, err := strconv.Atoi(roleIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Auth(c *gin.Context)] [strconv.Atoi(roleIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	data, err := logic2.RoleLogic{}.Auth(roleId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) Auth(c *gin.Context)] [logic.RoleLogic{}.Auth(roleId)] failed, err:", zap.Error(err))
		con.error(c, CodeRoleAuthAccessErr)
		return
	}
	con.success(c, data)
}

// DoAuth 给角色授权的接口
func (con RoleController) DoAuth(c *gin.Context) {
	// 解析参数
	p := new(models.DoAuthParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAuth(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeRoleAuthAccessErr)
		return
	}
	fmt.Println(p)
	// 业务逻辑
	err := logic2.RoleLogic{}.DoAuth(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAuth(c *gin.Context)] [logic.RoleLogic{}.DoAuth(p)] failed, err:", zap.Error(err))
		con.error(c, CodeRoleAuthAccessErr)
		return
	}
	con.success(c, true)
}
