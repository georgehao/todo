package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
)

type RegisterController struct {
	beego.Controller
}

func (this *RegisterController) Get()  {
	fmt.Println("register get")
	this.TplName = "register.html"
	this.Render()
}

func (this *RegisterController) Post()  {
	fmt.Println("register post")
}
