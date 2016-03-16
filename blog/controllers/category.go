package controllers

import (
	"github.com/astaxie/beego"
	"golang/blog/models"
)

type CategoryController struct {
	beego.Controller
}

func (c *CategoryController) Get() {
	op := c.Input().Get("op")

	switch op {
	case "add":
		name := c.Input().Get("name")
		if len(name) == 0 {
			break
		}
		err := models.AddCategory(name)
		if err != nil {
			beego.Error(err)
		}

		c.Redirect("/category", 301)
		return
	case "del":
		id := c.Input().Get("id")
		if len(id) == 0 {
			break
		}
		err := models.DelCategory(id)
		if err != nil {
			beego.Error(err)
		}
		c.Redirect("/category", 301)
		return
	}

	c.TplName = "category.html"

	c.Data["Title"] = "我的博客首页"
	c.Data["IsCategory"] = true

	var err error
	c.Data["Category"], err = models.GetAllCategorys()

	if err != nil {
		beego.Error(err)
	}
	return
}
