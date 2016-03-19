package controllers

import (
	"github.com/astaxie/beego"
	"golang/blog/models"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.TplName = "home.html"

	this.Data["Title"] = "我的博客首页"
	this.Data["IsHome"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(true)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}
