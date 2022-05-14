package tools

import (
	"math/rand"
	"strconv"
)

func GenerateCode() string {
	str := ""
	for i := 0; i < 4; i++ {
		str += strconv.Itoa(rand.Intn(10))
	}
	return str
}
