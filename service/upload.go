package service

import (
	"fmt"
	"gin-mall-tmp/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string)(filePath string, err error){
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.AvaterPath + "user" +bId + "/" //Belong to the user path
	if !DirExistOrNot(basePath){
		CreateDir(basePath)
	}
	avatarPath := basePath+userName+".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil{
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil{
		return
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

func UploadProductToLocakStatic(file multipart.File, userId uint, productName string)(filePath string, err error){
	bId := strconv.Itoa(int(userId))
	basePath := "." + conf.ProductPath + "boss" +bId + "/" //Belong to the user path
	if !DirExistOrNot(basePath){
		CreateDir(basePath)
	}
	productPath := basePath+productName+".jpg"
	fmt.Println("Saving file to:", productPath)
	content, err := ioutil.ReadAll(file)
	if err != nil{
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, 0666)
	if err != nil{
		return
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

//Judge whether exist or not
func DirExistOrNot(fileAddr string) bool {
	s,err := os.Stat(fileAddr)
	if err != nil{
		return false
	}
	return s.IsDir()
}
//Create file
func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil{
		return false
	}
	return true
}
