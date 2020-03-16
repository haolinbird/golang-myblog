package controllers

import (
	"golang-myblog/models"
)

type TagsController struct {
	BaseController
}

// 显示标签页
func (this *TagsController) Get() {
	// 获取所有文章的tag信息
	tags := models.QueryArticleWithParam("tags")

	// 设置模板数据
	this.Data["Tags"] = models.HandleTagsListData(tags)

	// 设置渲染模板
	this.TplName = "tags.html"
}
