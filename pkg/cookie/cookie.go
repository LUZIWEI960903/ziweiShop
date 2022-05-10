package cookie

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ginCookie struct {
}

func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	c.SetCookie(key, string(bytes), 3600, "/", "127.0.0.1:8080", false, true)
}

func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err1 := c.Cookie(key)
	if err1 == nil && valueStr != "" && valueStr != "[]" {
		err2 := json.Unmarshal([]byte(valueStr), obj)
		return err2 == nil
	}
	return false
}

var Cookie = &ginCookie{}
