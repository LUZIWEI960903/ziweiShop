package middlewares

import (
	"fmt"
	"net/http"
	"strings"
	"ziweiShop/controllers/admin"
	"ziweiShop/dao/mysql"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func InitAdminAuthMiddleware(c *gin.Context) {
	// 获取url
	pathname := strings.Split(c.Request.URL.String(), "?")[0]
	// 获取session
	session := sessions.Default(c)
	username, ok := session.Get("username").(string)
	roleId, _ := session.Get("roleId").(int)
	managerId, _ := session.Get("managerId").(int)
	isSuper, _ := session.Get("isSuper").(int)
	fmt.Printf("username:%v, roleId:%v, managerId:%v, isSuper:%v\n", username, roleId, managerId, isSuper)

	if !ok || username == "" {
		if pathname != "/admin/login" && pathname != "/admin/captcha" {
			c.Redirect(http.StatusFound, "/admin/login")
			return
		}
	} else { // 执行到这，代表登陆成功，需要判断该用户是否有权限访问其他 url
		fmt.Printf("Welcome %s~~\n", username)

		// 查询所有url，并存入urlMap
		accessUrlList := mysql.GetAllAccessUrl()
		urlMap := make(map[string]bool, len(accessUrlList))
		if accessUrlList != nil {
			for _, accessUrl := range accessUrlList {
				urlMap[accessUrl.Url] = true
			}
		}
		//fmt.Printf("urlMap:%v\n", urlMap)
		url := strings.Replace(pathname, "/admin", "", 1)
		if _, ok := urlMap[url]; ok && isSuper != 1 {
			// 根据 roleId 查询该 manager 对应的 role 其下的所有 access
			accessList, err := mysql.GetAccessListByRoleId(roleId)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"errcode": admin.CodeGetAccessErr,
					"errmsg":  admin.CodeGetAccessErr.Msg(),
				})
				return
			}
			// 把所有accessId 存放至map中
			accessListMap := make(map[int]int, len(accessList))
			for _, access := range accessList { // 遍历每个access对象
				accessListMap[access.AccessId] = access.AccessId
			}
			// 根据 url 在access表中查询 accessId
			accessId := mysql.GetAccessIdByUrl(url)
			if accessId == -1 {
				c.JSON(http.StatusOK, gin.H{
					"errcode": admin.CodeGetAccessErr,
					"errmsg":  admin.CodeGetAccessErr.Msg(),
				})
				return
			}
			// 判断 accessId 是否 在accessListMap中， 如果不是则代表无权访问该 url
			if _, ok := accessListMap[accessId]; !ok {
				c.JSON(http.StatusOK, gin.H{
					"errcode": admin.CodeInvalidAccess,
					"errmsg":  admin.CodeInvalidAccess.Msg(),
				})
				c.Abort()
			}
		}

	}
	fmt.Printf("Path: %s\n", pathname)
}
