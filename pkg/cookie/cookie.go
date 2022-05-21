package cookie

import (
	"encoding/json"
	"ziweiShop/pkg/tools"

	"github.com/gin-gonic/gin"
)

type ginCookie struct {
}

func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	// des加密
	encData, _ := tools.DesEncrypt(bytes)
	c.SetCookie(key, string(encData), 3600*24, "/", c.Request.Host, false, true)
}

func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		// dec解密
		decData, err2 := tools.DesDecrypt([]byte(valueStr))
		err3 := json.Unmarshal(decData, obj)
		return err2 == nil && err3 == nil
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}

var Cookie = &ginCookie{}
