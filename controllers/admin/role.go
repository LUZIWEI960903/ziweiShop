package admin

import (
	"ziweiShop/models"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	BaseController
}

// Index 显示管理员角色页面的接口
func (con RoleController) Index(c *gin.Context) {

}

// Add 给管理员增加角色页面的接口
func (con RoleController) Add(c *gin.Context) {

}

// DoAdd 给管理员增加角色的接口
func (con RoleController) DoAdd(c *gin.Context) {
	// 解析参数
	p := new(models.AddRoleParams)
	if err := c.ShouldBindJSON(p); err != nil {

		return
	}
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
