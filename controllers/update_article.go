package controllers

import (
	"golang-myblog/models"
)

type UpdateArticleController struct {
	BaseController
}

// 显示更新文章页面
func (this *UpdateArticleController) Get() {
	// 获取文章id
	id, _ := this.GetInt("id")

	// 获取文章信息
	article := models.GetArticleWithId(id)

	// 设置模板信息
	this.Data["Title"]   = article.Title
	this.Data["Tags"]    = article.Tags
	this.Data["Short"]   = article.Short
	this.Data["Content"] = article.Content
	this.Data["Id"]      = article.Id

	this.TplName = "write_article.html"
}

// 更新文章
func (this *UpdateArticleController) Post() {
	// 获取表单数据
	id, _ := this.GetInt("id")

	title := this.GetString("title")
	tags  := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	// 设置文章信息
	article := models.Article{id,title,tags,short,content,"", 0}

	// 修改文章信息
	_, err := models.UpdateArticle(article)
	if (err != nil) {
		this.Data["json"] = map[string]interface{}{"code": 10101, "message": "更新失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "更新成功"}
	}

	this.ServeJSON()
}

