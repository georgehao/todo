package routers

import (
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/controllers"
)

func init() {
	// login
	beego.Router("/login", &controllers.LoginController{}, "get,post:Login")
	beego.Router("/signup", &controllers.SignupController{}, "get,post:Signup")
	beego.Router("/logout", &controllers.LogoutController{}, "get,post:Logout")

	// index
	beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/add", &controllers.IndexController{}, "post:Add")
	beego.Router("/modify", &controllers.IndexController{}, "put:Modify")
	beego.Router("/about", &controllers.IndexController{}, "get:About")
}
