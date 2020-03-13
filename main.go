package main

import (
	_ "myblognew/routers"
	"github.com/astaxie/beego"
	"myblognew/utils"
)

func main() {
	utils.InitMysql()
	beego.Run()
}

