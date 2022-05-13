package captcha

import (
	"image/color"
	"ziweiShop/dao/redis"

	"github.com/mojocn/base64Captcha"
)

//var store = base64Captcha.DefaultMemStore
// 重写Get, Set, Verify方法，实现 base64Captcha 的Store接口，实现自定义的redisStore用于保存验证码
var store = redis.RedisStore{}

const SOURCEKEY = "1234567890qwertyuiopasdfghjklzxcvbnm"

// GenCaptcha 生成验证码
func GenCaptcha() (id string, b64s string, err error) {
	var driver base64Captcha.Driver
	driverString := base64Captcha.DriverString{
		Height:          40,
		Width:           100,
		NoiseCount:      0,
		ShowLineOptions: 2 | 4,
		Length:          4,
		Source:          SOURCEKEY,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 125,
		},
		Fonts: []string{"wqy-microhei.ttc"},
	}
	driver = driverString.ConvertFonts()

	c := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err = c.Generate()

	return id, b64s, err
}

// VerifyCaptcha 校验验证码
func VerifyCaptcha(id string, verifyValue string) bool {
	if store.Verify(id, verifyValue, true) {
		return true
	} else {
		return false
	}
}
