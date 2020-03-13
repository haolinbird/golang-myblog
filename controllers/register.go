package controllers

import (
	"github.com/astaxie/beego"
	"myblognew/models"
	"time"
	"fmt"
	"myblognew/utils"
)

type RegisterController struct {
	beego.Controller
}

/*
 * 渲染注册页面
 */
func (this *RegisterController) Get() {
	this.TplName = "register.html"
}

/*
 * 用户注册
 */
func (this *RegisterController) Post() {
	// 获取提交的表单信息
	username := this.GetString("username")
	password := this.GetString("password")
	repassword := this.GetString("repassword")

	// 先判断确认密码是否正确
	if password != repassword {
		// 定义返回json字符串,定义map的值类型为interface可以设置多种类型的值
		this.Data["json"] = map[string]interface{}{"code": 10001, "message": "密码不一致"}
		this.ServeJSON()
		return
	}

	// 先判断用户名是否存在
	id := models.QueryUserWithUsername(username)
	fmt.Println("username 对应的id:", id)
	if id > 0 {
		// 定义返回json字符串,定义map的值类型为interface可以设置多种类型的值
		this.Data["json"] = map[string]interface{}{"code": 10002, "message": "用户名已存在"}
		this.ServeJSON()
		return
	}

	// 加密明文密码
	password = utils.MD5(password)
	fmt.Println("md5后的密码:", password)

	   // ===注册用户===
	   // 定义结构体变量
	   user := models.User{Id: 0, Username: username, Password: password, Status: 0, Createtime: time.Now().Unix()}
	   // 注册用户
	   _, err := models.InsertUser(user)
	   if err != nil {
		this.Data["json"] = map[string]interface{}{"code": 10003, "message": "注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "注册成功"}
	}
	this.ServeJSON()
}


