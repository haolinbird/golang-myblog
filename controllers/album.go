package controllers

import (
	"github.com/Sirupsen/logrus"
	"golang-myblog/models"
)

type AlbumController struct {
	BaseController
}

// 显示相册页
func (this *AlbumController) Get() {
	// 查找所有上传的图片
	albums, err := models.FindAllAlbums()
	if err != nil {
		logrus.Error(err)
	}

	this.Data["Album"] = albums

	// 设置渲染模板
	this.TplName = "album.html"
}
