package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/lib"
	"html/template"
	"github.com/astaxie/beego/logs"
	//"time"
)

type LoginController struct {
	BaseController
}

func (this *LoginController) Login() {
	fmt.Println("login controller, method", this.Ctx.Input.Method())

	if this.IsLogin {
		this.Ctx.Redirect(302, this.URLFor("IndexController.Index"))
		return
	}

	this.TplName = "login.tpl"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	// 判断当前方法是否是post方法, 如果是Get方法,则直接返回, 呈现页面
	if !this.Ctx.Input.IsPost() {
		//u := models.User{Email:"haohongfan@ling.ai", Password:"123456", LastLoginTime:time.Now()}
		//fmt.Println(u.Email)
		//u.Create()
		fmt.Println("this method isn't post, return")
		return
	}

	flash := beego.NewFlash()
	email := this.GetString("email")
	password := this.GetString("password")

	// 判断用户名和密码是否有错误
	user, err := lib.Authenticate(email, password)
	if err != nil || user.ID < 1 {
		logs.Warning(err.Error())
		flash.Warning(err.Error())
		flash.Store(&this.Controller)
		return
	}

	flash.Success("Success Login in")
	flash.Store(&this.Controller)

	this.SetLogin(&user)
	this.Redirect(this.URLFor("IndexController.Index"), 303)
}

func (this *LoginController) Logout() {
	this.TplName = "login.tpl"
	this.Render()
}
