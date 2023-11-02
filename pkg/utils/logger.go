package utils

import (
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func init() {
	src, _ := setOutPutFile()
	if LogrusObj != nil {
		LogrusObj.Out = src
		return
	}
	//实例化
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{TimestampFormat: "2006-01-02 15:04:05"})
	LogrusObj = logger
}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	LogFilePath := ""
	dir, err := os.Getwd()
	//dir = filepath.Join(dir, "..")
	if err == nil {
		LogFilePath = dir + "/logs/"
	}
	_, err = os.Stat(LogFilePath)
	if os.IsNotExist(err) { //判断该文件路径是否存在
		err = os.MkdirAll(LogFilePath, 0777)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//日志文件名称
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(LogFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) { //判断该文件路径是否存在
		_, err = os.Create(fileName)
		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	//写入文件
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return nil, err
	}
	return src, nil
}
