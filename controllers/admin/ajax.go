package admin

import (
	"ziweiShop/logic"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AJAXController struct {
	BaseController
}

// ChangeStatus 使用 Ajax 更新指定表的指定id的status
func (con AJAXController) ChangeStatus(c *gin.Context) {
	// 解析参数
	id := c.Query("id")
	table := c.Query("table")
	field := c.Query("field")

	// 业务逻辑
	err := logic.AjaxLogic{}.ChangeStatus(id, table, field)
	if err != nil {
		zap.L().Error("[pkg: admin] [(con AJAXController) ChangeStatus(c *gin.Context)] [logic.AjaxLogic{}.ChangeStatus(id, table, field)] failed, err:", zap.Error(err))
		con.error(c, CodeAjaxChangeStatusErr)
		return
	}
	con.success(c, true)
}
