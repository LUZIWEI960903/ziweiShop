package tools

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "https://github.com/LUZIWEI960903/ziweiShop"

func MD5(s string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(s)))
}
