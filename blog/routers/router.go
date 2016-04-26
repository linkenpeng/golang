package routers

import (
	"golang/blog/controllers"
	"os"

	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func init() {
	i18n.SetMessage("zh-CH", "conf/local_zh-CN.ini")
	i18n.SetMessage("en-US", "conf/local_en-US.ini")
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/hello", &controllers.HelloController{})
	beego.Router("/category", &controllers.CategoryController{})
	beego.Router("/topic", &controllers.TopicController{})
	beego.Router("/lang", &controllers.LangController{})
	beego.AutoRouter(&controllers.TopicController{})
	beego.AddFuncMap("i18n", i18n.Tr)
	// 创建附件目录
	os.Mkdir("attachment", os.ModePerm)

	// 静态文件
	//beego.SetStaticPath("/attachment", "attachment")

	//作为单独控制处理
	beego.Router("/attachment/:all", &controllers.AttachController{})
}
