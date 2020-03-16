package controllers

type AboutMeController struct {
	BaseController
}

func (this *AboutMeController) Get() {
	this.Data["wechat"] = "微信：wx10000"
	this.Data["qq"] = "QQ：10000"
	this.Data["tel"] = "Tel：13688888888"
	this.TplName = "aboutme.html"
}