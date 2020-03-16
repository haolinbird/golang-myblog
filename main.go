package main

import (
	_ "golang-myblog/routers"
	"github.com/astaxie/beego"
	"golang-myblog/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

