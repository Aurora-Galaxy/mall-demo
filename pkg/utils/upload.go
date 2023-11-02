package utils

import (
	"gin_mall/conf"
	logging "github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadProductToLocalStatic 上传图片
func UploadProductToLocalStatic(file multipart.File, bossId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(bossId)) //转成string用于拼接路径
	//tempPath, _ := os.Getwd()        //获取当前工作目录
	//tempPath = filepath.Join(tempPath, "..")
	//basePath := tempPath + conf.Config.PhotoPath.ProductPath + "boss" + bId + "/"
	basePath := "." + conf.Config.PhotoPath.ProductPath + "boss" + bId + "/"
	err = DirExistOrNot(basePath)
	if err != nil {
		logging.Info(err)
		return "", err
	}
	productPath := basePath + productName + ".jpg"
	//content, err := ioutil.ReadAll(file)
	destFile, err := os.Create(productPath)
	defer destFile.Close()
	if err != nil {
		return "", err
	}
	_, err = io.Copy(destFile, file)
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// DirExistOrNot 判断文件是否存在,如果不存在直接创建
func DirExistOrNot(fileAddr string) error {
	_, err := os.Stat(fileAddr)
	if os.IsNotExist(err) {
		// 文件夹不存在，创建文件夹
		errDir := os.MkdirAll(fileAddr, 0755)
		if errDir != nil {
			logging.Info(errDir)
			return errDir
		}
	}
	return nil
}
