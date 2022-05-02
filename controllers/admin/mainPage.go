package admin

import (
	"fmt"
	logic "ziweiShop/logic/admin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type MainPageController struct {
	BaseController
}

// Index 显示后台主页面的接口
func (con MainPageController) Index(c *gin.Context) {
	session := sessions.Default(c)
	username, ok := session.Get("username").(string)
	if !ok {
		con.error(c, CodeNeedToLogin)
		return
	}
	fmt.Println("Welcome ", username)
	managerId, ok := session.Get("managerId").(int)
	if !ok {
		con.error(c, CodeNeedToLogin)
		return
	}
	//fmt.Printf("main page managerid%v managername:%v\n", managerId, username)
	// 业务逻辑
	data, err := logic.MainPageLogic{}.GetAdminIndexPageInfo(managerId)
	if err != nil {
		con.error(c, CodeGetAdminIndexPageInfoErr)
		return
	}
	con.success(c, data)
}
