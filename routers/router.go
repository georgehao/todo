package routers

import (
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/register", &controllers.RegisterController{})
}
