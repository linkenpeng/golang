package routers

import (
	"golang/blog/controllers"
	"os"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.AutoRouter(&controllers.TopicController{})

	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)
	// 静态文件
	beego.SetStaticPath("/attachment", "attachment")
}
