package tools

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func UploadImg(c *gin.Context, picName string) (string, error) {
	// 获取上传文件
	file, err := c.FormFile(picName)
	if err != nil {
		return "", err
	}
	// 判断后缀名是否正确
	extName := path.Ext(file.Filename)

	allowExtMap := map[string]bool{
		".jpg":  true,
		".JPG":  true,
		".png":  true,
		".PNG":  true,
		".gif":  true,
		".GIF":  true,
		".jpeg": true,
		".JPEG": true,
	}
	if _, ok := allowExtMap[extName]; !ok {
		return "", errors.New("Invalid file extName")
	}
	// 创建文件保存目录
	day := GetDate()
	dir := "./static/upload/" + day
	if err := os.MkdirAll(dir, 0666); err != nil {
		return "", err
	}
	fileName := fmt.Sprintf("%v%v", GetUnixN(), extName)
	dst := path.Join(dir, fileName)
	c.SaveUploadedFile(file, dst)
	return dst, nil
}
