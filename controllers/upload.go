package controllers

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"
	"golang-myblog/models"
)

type UploadController struct {
	BaseController
}

// 上传图片
func (this *UploadController) Post() {
	// 获取表单文件提交信息
    fileData, fileHeader, err := this.GetFile("upload")
    if err != nil {
		this.responseErr(err)
		return
	}

    // 获取当前时间
    now := time.Now()

    fileType := "other"
    // 判断文件后缀名
    fileExt := filepath.Ext(fileHeader.Filename)

    if fileExt == ".jpeg" || fileExt == ".gif" || fileExt == ".png" || fileExt == ".jpg" {
    	fileType = "img"
	}
	fileDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", fileType, now.Year(), now.Month(), now.Day())
	//创建存储目录，ModePerm是0777，这样拥有该文件夹路径的执行权限
	err = os.MkdirAll(fileDir, os.ModePerm)
	if err != nil {
		this.responseErr(err)
		return
	}

	//文件路径
	timeStamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d-%s", timeStamp, fileHeader.Filename)
	filePathStr := filepath.Join(fileDir, fileName)
	desFile, err := os.Create(filePathStr)
	if err != nil {
		this.responseErr(err)
		return
	}

	//将浏览器客户端上传的文件拷贝到本地路径的文件里面
	_, err = io.Copy(desFile, fileData)
	if err != nil {
		this.responseErr(err)
		return
	}

	if fileType == "img" {
		album := models.Album{0, filePathStr, fileName, 0, timeStamp}
		models.InsertAlbum(album)
	}

	this.Data["json"] = map[string]interface{}{"code": 1, "message": "上传成功"}
	this.ServeJSON()

}

// 返回错误json
func (this *UploadController) responseErr(err error) {
	this.Data["json"] = map[string]interface{}{"code": 0, "message": err}
	this.ServeJSON()
}

