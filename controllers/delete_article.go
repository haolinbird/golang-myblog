package controllers

import (
	"fmt"
	"golang-myblog/models"
	"log"
)

type DeleteArticleController struct {
	BaseController
}

// 删除文章后回到首页
func (this *DeleteArticleController) Get() {
	id, _  := this.GetInt("id")
	fmt.Print("删除文章id:", id)

	// 删除文章
	_, err := models.DeleteArticle(id)
	if err != nil {
		log.Println(err)
	}

	// 删除完成后重定向回到首页
	this.Redirect("/", 302)
}
