package routers

import (
	"github.com/astaxie/beego"
	"svnmanage/controllers"
)

func init() { //注册静态路由
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})
	beego.Router("/svn", &controllers.SvnController{})
	beego.Router("/exec", &controllers.ExecController{})
	beego.Router("/logout", &controllers.LogoutController{})
	beego.Router("/edit", &controllers.EditController{})
	beego.Router("/add", &controllers.AddController{})
	beego.SetStaticPath("/log", "log")
}
