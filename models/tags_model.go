package models

import "strings"

func HandleTagsListData(tags []string) map[string]int {
	// 初始化map变量
	var tagsMap = make(map[string]int)

	// 计算标签对应的文章数量
	for _, tag := range tags {
        tagList := strings.Split(tag, "&")
        for _, value := range  tagList {
        	tagsMap[value]++
		}
	}

	return tagsMap
}
