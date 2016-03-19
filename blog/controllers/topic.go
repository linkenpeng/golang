package controllers

import (
	"github.com/astaxie/beego"
	"golang/blog/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) Get() {
	this.TplName = "topic.html"
	this.Data["Title"] = "文章列表"
	this.Data["IsTopic"] = true
	this.Data["IsLogin"] = checkAccount(this.Ctx)

	topics, err := models.GetAllTopics(false)
	if err != nil {
		beego.Error(err)
	} else {
		this.Data["Topics"] = topics
	}
}

func (this *TopicController) Post() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	beego.Error(tid)
	var err error
	if len(tid) == 0 {
		err = models.AddTopic(title, content)
	} else {
		err = models.ModifyTopic(tid, title, content)
	}

	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}

func (this *TopicController) Add() {
	this.Data["Title"] = "添加文章"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic_add.html"
}

func (this *TopicController) View() {
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.Data["IsTopic"] = true
	this.TplName = "topic_view.html"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
		return
	}

	this.Data["Topic"] = topic
	this.Data["Title"] = topic.Title
	this.Data["Tid"] = tid
}

func (this *TopicController) Modify() {
	this.TplName = "topic_add.html"
	this.Data["Title"] = "修改文章"

	tid := this.Input().Get("tid")
	topic, err := models.GetTopic(tid)
	if err != nil {
		beego.Error(err)
		this.Redirect("/", 302)
	}

	this.Data["Topic"] = topic
	this.Data["Title"] = topic.Title
	this.Data["Tid"] = tid
}

func (this *TopicController) Delete() {
	if !checkAccount(this.Ctx) {
		this.Redirect("/login", 302)
		return
	}

	tid := this.Input().Get("tid")
	err := models.DeleteTopic(tid)
	if err != nil {
		beego.Error(err)
	}

	this.Redirect("/topic", 302)
}
