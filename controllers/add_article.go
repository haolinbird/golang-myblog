package controllers

import (
	"github.com/astaxie/beego"
	"golang-myblog/models"
	"time"
)

type AddArticleController struct {
	beego.Controller
}

// 显示写文章页面
func (this *AddArticleController) Get() {
    this.TplName = "write_article.html"
}

// 添加文章
func (this *AddArticleController) Post() {
	// 获取表单数据
	title := this.GetString("title")
	tags  := this.GetString("tags")
	short := this.GetString("short")
	content := this.GetString("content")

	// 设置文章信息
	article := models.Article{0,title,tags,short,content,"haolin", time.Now().Unix()}

	// 添加文章
	_, err := models.AddArticle(article)
	if (err != nil) {
        this.Data["json"] = map[string]interface{}{"code": 10101, "message": "error"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "ok"}
	}

	this.ServeJSON()
}
