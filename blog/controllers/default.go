package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "home.html"

	c.Data["Title"] = "我的博客首页"
	c.Data["IsHome"] = true

	c.Data["IsLogin"] = checkAccount(c.Ctx)
	return
}
