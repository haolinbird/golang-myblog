package controllers

import (
	"golang-myblog/models"
)

type HomeController struct {
	BaseController
}

/*
 * 渲染首页
 */
func (this *HomeController) Get() {
	// 获取标签信息
	tag := this.GetString("tag")
	// 设置当前分页
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	// 定义一个切片变量存储文章信息
	var articleList []models.Article

    if len(tag) > 0 {
        // 根据标签搜索文章
		articleList, _ = models.GetArticleListWithTag(tag)
        this.Data["HasFooter"] = false
	} else {
		// 根据分页获取文章列表
		articleList, _ = models.GetArticleListWithPage(page, 10)
		this.Data["PageCode"] = models.ConfigHomeFooterPageCode(page)
		this.Data["HasFooter"] = true
	}

	this.Data["Content"] = models.MakeHomeBlocks(articleList, this.IsLogin)

	this.TplName = "home.html"
}
