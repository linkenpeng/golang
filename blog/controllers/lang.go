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
	//beego.Info(lang)
	if lang == "zh-CN" {
		this.Lang = "zh-CN"
	} else {
		this.Lang = "en-US"
	}
	this.Data["Lang"] = this.Lang
}

type LangController struct {
	LangBaseController
}

func (this *LangController) Get() {
	this.Data["Hi"] = this.Tr("hi")
	this.Data["Bye"] = this.Tr("bye")
	this.TplName = "lang.tpl"
}
