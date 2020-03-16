package models

import (
	"golang-myblog/utils"
	"fmt"
	"log"
	"strconv"
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

	result,err := utils.ModifyDB(sql, article.Title, article.Author, article.Tags, article.Short, article.Content, article.Createtime)
	// 添加成功后更新文章总数
	if err == nil {
		SetArticleRowsNum()
	}

    return result, err
}

// 获取文章列表
func GetArticleListWithPage(page int, pagesize int) ([]Article, error) {
    // 构造查询数据
	page--
    sql := fmt.Sprintf("select id,title,author,tags,short,content,createtime from article limit %d,%d", page * pagesize, pagesize)

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

// 获取指定文章信息
func GetArticleWithId(id int) Article {
	row := utils.QueryRowDB("select id,title,tags,short,content,author,createtime from article where id=" + strconv.Itoa(id))

	// 读取文章数据
	title   := ""
	tags    := ""
	short   := ""
	content := ""
	author  := ""
	var createtime int64
	createtime = 0
	row.Scan(&id, &title, &tags, &short, &content, &author, &createtime)

	// 返回数据
	art := Article{id, title, tags, short, content, author, createtime}

	return art
}

// 更新文章信息
func UpdateArticle(article Article) (int64, error) {
    return utils.ModifyDB("update article set title=?,tags=?,short=?,content=? where id = ?", article.Title, article.Tags, article.Short, article.Content, article.Id)
}

// 删除文章
func DeleteArticle(id int) (int64, error) {
	// 删除文章
	result, err := utils.ModifyDB("delete from article where id = ?", id)

	// 删除文章成功后更新文章总数
	if err == nil {
		SetArticleRowsNum()
	}

	return result, err
}

// 查询标签，返回一个字段的列表
func QueryArticleWithParam(param string) []string {
	rows, err := utils.QueryDB(fmt.Sprintf("select %s from article", param))
	if err != nil {
		log.Println(err)
	}

	var paramList []string
	for rows.Next() {
		arg := ""
		rows.Scan(&arg)
		paramList = append(paramList, arg)
	}

	return paramList
}

// 根据标签获取文章列表
func GetArticleListWithTag(tag string) ([]Article, error) {
	// 构造查询数据
	sql := "select id,title,author,tags,short,content,createtime from article"
	sql += " where tags like '%&" + tag + "&%'"
	sql += " or tags like '%&" + tag + "'"
	sql += " or tags like '" + tag + "&%'"
	sql += " or tags like '" + tag + "'"

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