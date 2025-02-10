package util

import (
	"log"
	"os"
	"path"
	"time"
	"github.com/sirupsen/logrus"
)

var LogrusObj *logrus.Logger

func init(){
	src, _ := setOutPutFile()
	if LogrusObj != nil{
		LogrusObj.Out = src
		return
	}
	//Instance
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2025-01-01 12:00:00",
	})
	LogrusObj = logger
}

func setOutPutFile() (*os.File, error){
	now := time.Now()
	logFilePath := ""
	if idr, err := os.Getwd();err == nil{ //Require work dectiory
		logFilePath = idr + "/logs/"
	}
	_,err := os.Stat(logFilePath)
	if os.IsNotExist(err){
		if err = os.MkdirAll(logFilePath, 0777); err != nil{
			log.Println(err.Error())
			return nil, err
		}
	}
	logFileName := now.Format("2025-01-01") + ".log"

	fileName := path.Join(logFilePath, logFileName)
	if os.IsNotExist(err){
		if err = os.MkdirAll(fileName, 0777);err != nil{
			log.Println(err.Error())
			return nil, err
		}
	}
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil{
		return nil, err
	}
	return src, nil

}