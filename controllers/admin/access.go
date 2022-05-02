package admin

import (
	"strconv"
	"ziweiShop/logic"
	"ziweiShop/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type AccessController struct {
	BaseController
}

// Index 显示权限列表页面的接口
func (con AccessController) Index(c *gin.Context) {
	// 业务逻辑
	topAccessListWithAccessList, err := logic.AccessLogic{}.GetTopAccessListWithAccessList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Index(c *gin.Context)] [logic.AccessLogic{}.GetTopAccessListWithAccessList()] failed, err:", zap.Error(err))
		con.error(c, CodeTopAccessListWithAccessListErr)
		return
	}
	con.success(c, topAccessListWithAccessList)
}

// Add 显示增加权限页面的接口
func (con AccessController) Add(c *gin.Context) {
	// 业务逻辑
	topAccessList, err := logic.AccessLogic{}.GetTopAccessList()
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Add(c *gin.Context)] [logic.AccessLogic{}.GetTopAccessList()] failed, err:", zap.Error(err))
		con.error(c, CodeGetTopAccessListErr)
		return
	}
	con.success(c, topAccessList)
}

// DoAdd 执行增加权限的接口
func (con AccessController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddAccessParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 参数校验
	if err := verifyAddAccessParams(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [verifyAddAccessParams(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEmptyModuleName)
		return
	}
	// 业务逻辑
	err := logic.AccessLogic{}.DoAdd(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoAdd(c *gin.Context)] [logic.AccessLogic{}.DoAdd(p)] failed, err:", zap.Error(err))
		con.error(c, CodeAddAccessErr)
		return
	}
	con.success(c, true)
}

// Edit 编辑权限页面的接口
func (con AccessController) Edit(c *gin.Context) {
	// 解析参数
	accessIdStr := c.Query("id")
	accessId, err := strconv.Atoi(accessIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Edit(c *gin.Context)] [strconv.Atoi(accessIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	editAccess, err := logic.AccessLogic{}.GetAccessById(accessId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Edit(c *gin.Context)] [logic.AccessLogic{}.GetAccessById(accessId)] failed, err:", zap.Error(err))
		con.error(c, CodeGetAccessErr)
		return
	}
	con.success(c, editAccess)
}

// DoEdit 编辑权限的接口
func (con AccessController) DoEdit(c *gin.Context) {
	// 解析参数
	p := new(models.EditAccessParams)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoEdit(c *gin.Context)] [c.ShouldBindJSON(p)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 参数校验
	if err := verifyEditAccessParams(p); err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoEdit(c *gin.Context)] [verifyEditAccessParams(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEmptyModuleName)
		return
	}
	// 业务逻辑
	err := logic.AccessLogic{}.DoEdit(p)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) DoEdit(c *gin.Context)] [logic.AccessLogic{}.DoEdit(p)] failed, err:", zap.Error(err))
		con.error(c, CodeEditAccessErr)
		return
	}
	con.success(c, true)
}

// Delete 删除权限的接口
func (con AccessController) Delete(c *gin.Context) {
	// 解析参数
	accessIdStr := c.Query("id")
	accessId, err := strconv.Atoi(accessIdStr)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Delete(c *gin.Context)] [strconv.Atoi(accessIdStr)] failed, err:", zap.Error(err))
		con.error(c, CodeInValidParams)
		return
	}
	// 业务逻辑
	err = logic.AccessLogic{}.Delete(accessId)
	if err != nil {
		zap.L().Error("[pkg: admin] [func: (con AccessController) Delete(c *gin.Context)] [logic.AccessLogic{}.Delete(accessId))] failed, err:", zap.Error(err))
		con.error(c, CodeDeleteAccessErr)
		return
	}
	con.success(c, true)
}
