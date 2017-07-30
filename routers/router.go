package routers

import (
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/controllers"
)

func init() {
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/signup", &controllers.SignupController{}, "get,post:Signup")
	beego.Router("/logout", &controllers.LogoutController{}, "get,post:Logout")
}
