package models

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"strconv"
	"strings"
	"golang-myblog/utils"
)

type TagLink struct {
	TagName string
	TagUrl string
}

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string

	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

// 定义分页结构体
type HomeFooterPageCode struct {
	HasPre bool
	HasNext bool
	ShowPage string
	PreLink string
	NextLink string
}

// 显示首页内容
func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
    htmlHome := ""
    for _, article := range articles {
    	// 初始化一个首页内容结构体变量
    	homeParam := HomeBlockParam{}

    	// 填充数据
    	homeParam.Id = article.Id
    	homeParam.Title = article.Title
    	homeParam.Tags = createTagsLinks(article.Tags)
    	homeParam.Short = article.Short
    	homeParam.Content = article.Content
    	homeParam.Author = article.Author
    	homeParam.CreateTime = utils.TimestampToDate(article.Createtime)
    	homeParam.Link = "/article/" + strconv.Itoa(article.Id)
		homeParam.UpdateLink = "/article/update?id=" + strconv.Itoa(article.Id)
		homeParam.DeleteLink = "article/delete?id=" + strconv.Itoa(article.Id)
		homeParam.IsLogin = isLogin

		// 处理变量
		// ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}

		//就是将html文件里面的占位符替换为传进去的数据
		t.Execute(&buffer, homeParam)

		htmlHome += buffer.String()
	}

    return template.HTML(htmlHome)
}

// 将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsParam := strings.Split(tags, "&")
	for _, tag := range tagsParam {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}

	return tagLink
}


// 翻页
func ConfigHomeFooterPageCode(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	// 查询出总的条数
	num := GetArticleRowsNum()
	// 从配置文件读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	// 计算出总页数
	allPageNum := (num - 1) / pageRow + 1

	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)

	// 当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1{
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}

	// 当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}

	pageCode.PreLink = "/?page=" + strconv.Itoa(page - 1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page + 1)

	return pageCode
}

// 存储文章记录数，当文章新增或者删除时需要更新这个值
var artcileRowsNum = 0

func GetArticleRowsNum() int {
	if artcileRowsNum == 0 {
		artcileRowsNum = QueryArticleRowNum()
	}
	return artcileRowsNum
}

// 查询文章的总条数
func QueryArticleRowNum() int {
	row := utils.QueryRowDB("select count(id) from article")
	num := 0
	row.Scan(&num)
	return num
}

// 设置页数
func SetArticleRowsNum() {
	artcileRowsNum = QueryArticleRowNum()
}