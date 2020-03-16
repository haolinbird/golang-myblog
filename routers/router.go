package routers

import (
	"golang-myblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // 首页
    beego.Router("/", &controllers.HomeController{})

    // 注册功能
	beego.Router("/register", &controllers.RegisterController{})

    // 登录功能
    beego.Router("/login", &controllers.LoginController{})

    // 退出功能
    beego.Router("/exit", &controllers.LogoutController{})

    // 写文章
    beego.Router("/article/add", &controllers.AddArticleController{})

    // 显示文章内容
    beego.Router("/article/:id", &controllers.ShowArticleController{})
}
