package controllers

import (
	"myblognew/models"
)

type HomeController struct {
	BaseController
}

/*
 * 渲染首页
 */
func (this *HomeController) Get() {
	// 设置当前分页
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}

	this.Data["PageCode"] = 1
	this.Data["HasFooter"] = false

	// 定义一个切片变量存储文章信息
	var articleList []models.Article
	// 根据分页获取文章列表
	articleList, _ = models.GetArticleListWithPage(page, 10)
	this.Data["Content"] = models.MakeHomeBlocks(articleList, this.IsLogin)


	this.TplName = "home.html"
}
