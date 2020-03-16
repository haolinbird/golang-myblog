package controllers

import (
	"golang-myblog/models"
	"golang-myblog/utils"
	"strconv"
)

type ShowArticleController struct {
	BaseController
}

// 显示文章详情
func (this *ShowArticleController) Get() {
	// 获取GET请求传入的指定参数
	idStr := this.Ctx.Input.Param(":id")

	// 将输入参数转化为整型
	id, _ := strconv.Atoi(idStr)

    // 获取id对应的文章信息
    art := models.GetArticleWithId(id)

    // 设置模板数据
    this.Data["Title"] = art.Title
    this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)

    // 设置渲染模板
    this.TplName = "show_article.html"

}
