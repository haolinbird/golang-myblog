package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"myblognew/utils"
	"myblognew/models"
)

type LoginController struct {
	beego.Controller
}

/*
 * 渲染注册页面
 */
func (this *LoginController) Get() {
	this.TplName = "login.html"
}

/*
 * 用户登录
 */
func (this *LoginController) Post() {
    //获取表单数据
    username := this.GetString("username")
    password := this.GetString("password")

    // 打印调试信息
    fmt.Println("username: ", username, ",password: ", password)

    // 验证用户名和密码
    id := models.QueryUserWithUsernamePwd(username, utils.MD5(password))
    if id > 0 {
    	// 登录成功后将登录信息存储到session
    	this.SetSession("loginuser", username)

    	this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 10011, "message": "用户或者密码不正确"}
	}
	this.ServeJSON()
}


