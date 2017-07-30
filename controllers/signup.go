package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
	"github.com/georgehao/todo/models"
	"github.com/georgehao/todo/lib"
	"fmt"
)

type SignupController struct {
	BaseController
}

func (this *SignupController) Signup() {
	this.TplName = "signup.html"
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	if !this.Ctx.Input.IsPost() {
		return
	}

	flash := beego.NewFlash()
	u := &models.User{}
	var err error
	if err = this.ParseForm(u); err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		return
	}

	if err = models.IsValid(u); err != nil {
		flash.Error(err.Error())
		flash.Store(&this.Controller)
		return
	}

	id, err := lib.SignupUser(u)
	if err != nil || id < 1 {
		flash.Warning(err.Error())
		flash.Store(&this.Controller)
		fmt.Println("singup err:", err.Error())
		return
	}

	flash.Success("Register user")
	flash.Store(&this.Controller)

	this.SetLogin(u)
	this.Redirect(this.URLFor("IndexController.Index"), 303)
}
