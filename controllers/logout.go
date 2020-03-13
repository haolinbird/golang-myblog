package controllers

import (
	"github.com/astaxie/beego"
)

type LogoutController struct {
	beego.Controller
}

/*
 * 用户退出
 */
func (this *LogoutController) Get() {
	//清除该用户登录状态的数据
    this.DelSession("loginuser")
    //跳会到首页
    this.Redirect("/",302)
}


