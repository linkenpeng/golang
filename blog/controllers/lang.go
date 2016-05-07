package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

type LangBaseController struct {
	beego.Controller
	i18n.Locale
}

func (this *LangBaseController) Prepare() {
	lang := this.GetString("lang")
	if lang == "zh-CN" {
		this.Lang = lang
	} else {
		this.Lang = "en-US"
	}
	beego.Info(this.Lang)
	this.Data["Lang"] = this.Lang
}

type LangController struct {
	LangBaseController
}

func (this *LangController) Get() {
	this.Data["Hi"] = this.Tr("hi")
	this.Data["Bye"] = this.Tr("bye")
	//this.Data["About"] = this.Tr("about")

	//this.Data["Hi"] = "hi"
	//this.Data["Bye"] = "bye"
	//this.Data["About"] = "about"

	this.TplName = "lang.tpl"
}
