package admin

import (
	"errors"
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type ManagerController struct {
	BaseController
}

// Index 显示管理员列表页面的接口
func (con ManagerController) Index(c *gin.Context) {
	// 业务逻辑
	indexManagerList, err := logic.ManagerLogic{}.GetIndexManagerList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Index(c *gin.Context)] [logic.ManagerLogic{}.GetIndexManagerList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetIndexManagerListErr)
		return
	}
	con.success(c, indexManagerList)
}

// Add 显示增加管理员页面的接口
func (con ManagerController) Add(c *gin.Context) {
	// 返回roleList
	var managerService = logic.ManagerLogic{}
	roleList, err := managerService.GetRoleList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Add(c *gin.Context)] [managerService.GetRoleList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetRoleListErr)
		return
	}
	con.success(c, roleList)
}

// DoAdd 执行增加管理员的接口
func (con ManagerController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddManagerParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err := logic.ManagerLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoAdd(c *gin.Context)] [logic.ManagerLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		if errors.Is(err, logic.ErrorManagerExist) {
			con.error(c, CodeManagerExistErr)
			return
		}
		if errors.Is(err, logic.ErrorRoleNotExist) {
			con.error(c, CodeRoleNotExistErr)
			return
		}
		con.error(c, CodeAddManagerErr)
		return
	}
	con.success(c, true)
}

// Edit 显示修改管理员页面的接口
func (con ManagerController) Edit(c *gin.Context) {
	// 解析参数
	managerIdStr := c.Query("id")
	managerId, err := strconv.Atoi(managerIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Edit(c *gin.Context)] [strconv.Atoi(managerIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	var editManagerService = logic.ManagerLogic{}
	editManagerInfo, err := editManagerService.GetEditManagerInfo(managerId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Edit(c *gin.Context)] [editManagerService.GetEditManagerList(managerId)] failed, err:", zap.Error(err))
		con.error(c, CodeGetEditManagerErr)
		return
	}
	con.success(c, editManagerInfo)
}

// DoEdit 显示修改管理员信息的接口
func (con ManagerController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditManagerParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err := logic.ManagerLogic{}.DoEdit(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) DoEdit(c *gin.Context)] [logic.ManagerLogic{}.DoEdit(p) ] failed, err:", zap.Error(err))
		if errors.Is(err, logic.ErrorManagerNotExist) {
			// 输入的id不存在
			con.error(c, CodeManagerNotExist)
			return
		}
		if errors.Is(err, logic.ErrorManagerExist) {
			// 输入的roleId不存在
			con.error(c, CodeRoleNotExistErr)
			return
		}
		con.error(c, CodeManagerDoEditErr)
		return
	}
	con.success(c, true)
}

// Delete 删除管理员的接口
func (con ManagerController) Delete(c *gin.Context) {
	// 解析参数
	managerIdStr := c.Query("id")
	managerId, err := strconv.Atoi(managerIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Delete(c *gin.Context)] [strconv.Atoi(managerIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err1 := logic.ManagerLogic{}.DeleteManager(managerId)
	if err1 != nil {
		zap.L().Error("[pkg: admin] [func: (con ManagerController) Delete(c *gin.Context)] [logic.ManagerLogic{}.DeleteManager(managerId)] failed, err:", zap.Error(err))
		con.error(c, CodeDeleteManagerErr)
		return
	}
	con.success(c, true)
}
