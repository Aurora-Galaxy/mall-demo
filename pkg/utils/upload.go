package utils

import (
	"gin_mall/conf"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadProductToLocalStatic 上传图片
func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	strId := strconv.Itoa(int(userId)) //转成string用于拼接路径
	basePath := "." + conf.Config.PhotoPath.ProductPath + "product" + strId + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	create, err := os.Create(productPath)
	if err != nil {
		return "", err
	}
	_, err = io.Copy(create, file) //使用copy替代ioutil.readall
	if err != nil {
		return "", err
	}
	return "user" + strId + "/" + productName + ".jpg", nil
}

// DirExistOrNot 判断文件是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		log.Println(err)
		return false
	}
	return s.IsDir() //判断路径是否是一个目录
}

// CreateDir 创建文件夹
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 7550)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
