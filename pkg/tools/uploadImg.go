package tools

import (
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"path"
	"ziweiShop/config"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

const (
	ossAccessKey       = "LTAI5tPdVw4WUtxhpod2e9ia"
	ossAccessKeySecret = "ShvrPldc8p4pkYN7aC4NXTYhzJBhsZ"
	ossEndPoint        = "oss-cn-guangzhou.aliyuncs.com"
	ossBucketName      = "ziwei-shop"
)

func GetOssStatus() int64 {
	return config.Conf.OssStatus
}

func UploadImg(c *gin.Context, picName string) (string, error) {
	if GetOssStatus() == 1 {
		return OssUploadImg(c, picName)
	}
	return LocalUploadImg(c, picName)
}

// LocalUploadImg 上传图片到本地服务器
func LocalUploadImg(c *gin.Context, picName string) (string, error) {
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

// OssUpload 上传oss云服务器的具体执行方法
func OssUpload(file *multipart.FileHeader, dst string) (string, error) {
	client, err := oss.New(ossEndPoint, ossAccessKey, ossAccessKeySecret)
	if err != nil {
		return "", err
	}

	bucket, err := client.Bucket(ossBucketName)
	if err != nil {
		return "", err
	}

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	err = bucket.PutObject(dst, src)
	if err != nil {
		return "", err
	}

	return dst, nil
}

// OssUploadImg 上传图片到oss云服务器
func OssUploadImg(c *gin.Context, picName string) (string, error) {
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
	dir := "static/upload/" + day

	fileName := fmt.Sprintf("%v%v", GetUnixN(), extName)
	dst := path.Join(dir, fileName)

	return OssUpload(file, dst)
}
