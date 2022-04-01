package admin

import (
	"errors"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

// Index 显示管理员角色页面的接口
func (con RoleController) Index(c *gin.Context) {
	var roleService = logic.RoleLogic{}
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
	err := logic.RoleLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con RoleController) DoAdd(c *gin.Context)] [logic.RoleLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		if errors.Is(err, logic.ErrorRoleExist) {
			con.error(c, CodeRoleExist)
			return
		}
		if errors.Is(err, logic.ErrorAddRole) {
			con.error(c, CodeAddRoleErr)
			return
		}
	}
	con.success(c, true)
}

// Edit 给管理员编辑角色页面的接口
func (con RoleController) Edit(c *gin.Context) {

}

// DoEdit 给管理员编辑角色的接口
func (con RoleController) DoEdit(c *gin.Context) {

}

// Delete 给管理员删除角色的接口
func (con RoleController) Delete(c *gin.Context) {

}
