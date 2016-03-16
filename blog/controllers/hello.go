package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (c *HelloController) Get() {
	c.TplName = "hello.tpl"

	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"

	c.Data["Truecond"] = true

	type u struct {
		Name string
		Age  int
		Sex  string
	}

	user := &u{
		Name: "joe",
		Age:  20,
		Sex:  "Male",
	}
	c.Data["User"] = user

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	c.Data["Nums"] = nums

	c.Data["TplVar"] = "hey guys"

	c.Data["Html"] = "<div>Hello html.</div>"

	c.Data["Pipe"] = "<div>Hello beego.</div>"
}
