package models

import (
    "fmt"
    "golang-myblog/utils"
)

type User struct {
    Id         int
    Username   string
    Password   string
    Status     int
    Createtime int64
}

// 查询用户
func QueryUserWithUsername(username string) int {
	// 定义查询语句
    sql := fmt.Sprintf("select id from users where username = '%s'", username)
    fmt.Println(sql)
    // 查询一行记录
    row := utils.QueryRowDB(sql)
    // 获取记录对应字段的值
    id := 0
    row.Scan(&id)
    // 返回数据
    return id
}

// 验证用户名和密码
func QueryUserWithUsernamePwd(username string, password string) int {
    // 定义查询语句
    sql := fmt.Sprintf("select id from users where username = '%s' and password = '%s'", username, password)
    fmt.Println(sql)
    // 查询一行记录
    row := utils.QueryRowDB(sql)
    // 获取记录对应字段的值
    id := 0
    row.Scan(&id)
    // 返回数据
    return id
}

// 添加用户
func InsertUser(user User) (int64, error) {
    sql := "insert into users(username, password, status, createtime) values (?,?,?,?)"

    return utils.ModifyDB(sql, user.Username, user.Password, user.Status, user.Createtime)
}


