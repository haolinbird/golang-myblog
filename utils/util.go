package utils

import (
	"fmt"
	"crypto/md5"
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

