package routers

import (
	"github.com/astaxie/beego"
	"golang/blog/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/category", &controllers.CategoryController{})
}
