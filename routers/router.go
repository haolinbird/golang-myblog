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

    // 更新文章
    beego.Router("/article/update", &controllers.UpdateArticleController{})

    // 删除文章
    beego.Router("article/delete", &controllers.DeleteArticleController{})

    // 标签页
    beego.Router("/tags", &controllers.TagsController{})

    // 相册
    beego.Router("/album", &controllers.AlbumController{})

    // 上传图片
    beego.Router("/upload", &controllers.UploadController{})

    // 关于我们
    beego.Router("/aboutme", &controllers.AboutMeController{})
}
