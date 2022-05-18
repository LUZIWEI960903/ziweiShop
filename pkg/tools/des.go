package tools

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

var DESKEY = []byte("66666666")

//Des 加密
func DesEncrypt(origData []byte) ([]byte, error) {
	block, err := des.NewCipher(DESKEY)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, DESKEY)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

//Des 解密
func DesDecrypt(crypted []byte) ([]byte, error) {
	block, err := des.NewCipher(DESKEY)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, DESKEY)
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
