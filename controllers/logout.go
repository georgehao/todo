package controllers

import "github.com/astaxie/beego"

type LogoutController struct {
	beego.Controller
}

func (this *LogoutController) Logout() {
	flash := beego.NewFlash()
	flash.Success("Success logged out")
	flash.Store(&this.Controller)

	this.Ctx.Redirect(302, this.URLFor("LoginController.Login"))
}