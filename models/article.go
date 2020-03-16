package models

import (
	"golang-myblog/utils"
	"fmt"
)

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime int64
}

// 添加文章
func AddArticle(article Article) (int64, error) {
	sql := "insert into article(title,author,tags,short,content,createtime) values(?,?,?,?,?,?)"

	result,err := utils.ModifyDB(sql, article.Title, article.Tags, article.Short, article.Content, article.Author, article.Createtime)
	// 添加成功后更新文章总数
	if err == nil {
		SetArticleRowsNum()
	}

    return result, err
}

// 获取文章列表
func GetArticleListWithPage(page int, pagesize int) ([]Article, error) {
    // 构造查询数据
    sql := fmt.Sprintf("select id,title,author,tags,short,content,createtime from article limit %d,%d", page, pagesize)

    // 查询数据
    rows, err := utils.QueryDB(sql)
    if err != nil {
    	return nil, err
	}

    // 初始化一个切片来存储返回的文章列表
    var articleList []Article
    // 循环查询结果填充数据
    for rows.Next() {
    	id := 0
    	title := ""
    	tags := ""
    	short := ""
    	content := ""
    	author := ""
    	var createtime int64
    	createtime = 0
    	// 获取字段数据
    	rows.Scan(&id, &title, &tags, &short, &content, &author, &createtime)
        article := Article{id, title,tags,short,content,author,createtime}
        // 填充切片
        articleList = append(articleList, article)
	}

    return articleList, nil
}
