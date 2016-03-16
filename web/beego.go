package main

/*
官网：http://beego.me/
安装：go get github.com/astaxie/beego
更新：go get -u github.com/astaxie/beego
*/

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}

func (this *HomeController) Get() {
	this.Ctx.WriteString("Hello World!")
}

func runBeego() {
	beego.Router("/", &HomeController{})
	beego.Run()
}
