package utils

import (
	"bytes"
	"fmt"
	"crypto/md5"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"github.com/sourcegraph/syntaxhighlight"
	"html/template"
	"time"
)

/*
   md5加密
 */
func MD5(str string) string{
    md5str := fmt.Sprintf("%x", md5.Sum([]byte(str)))

    return md5str
}

// 时间戳转日期
func TimestampToDate(timestamp int64) string{
	// 转化所需模板
	timeLayout := "2006-01-02 15:04:05"
	return time.Unix(timestamp, 0).Format(timeLayout)
}

// 将markdown语法内容转换为html
func SwitchMarkdownToHtml(content string) template.HTML {
	// 获取markdown格式内容
	markdown := blackfriday.MarkdownCommon([]byte(content))

	// 转换为html文档
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader(markdown))

	/*
	  对document进程查询，选择器和css的语法一样
	  第一个参数：i是查询到的第几个元素
	  第二个参数：selection就是查询到的元素
	 */
	doc.Find("code").Each(func(i int, selection *goquery.Selection) {
		light, _ := syntaxhighlight.AsHTML([]byte(selection.Text()))
		selection.SetHtml(string(light))
		fmt.Println(selection.Html())
		fmt.Println("light:", string(light))
		fmt.Println("\n\n\n")
	})
	htmlString, _ := doc.Html()
	return template.HTML(htmlString)
}