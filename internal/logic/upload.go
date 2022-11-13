package logic

import (
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"wenba/internal/global"
)

func UploadAvatarToLocalStatic(file multipart.File, userID int64, userName string) (filePath string, err error) {
	//todo
	basePath := "../" + global.Settings.Rule.Avatar + "user" + strconv.FormatInt(userID, 10) + "/"
	fmt.Println(basePath)
	if DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return
	}
	return "user" + strconv.FormatInt(userID, 10) + "/" + userName + ".jpg", nil
}

//DirExistOrNot 判断文件夹路径是否存在
func DirExistOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func CreateDir(dirName string) bool {
	err := os.MkdirAll(dirName, 755)
	if err != nil {
		return false
	}
	return true
}
