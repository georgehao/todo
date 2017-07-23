package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/models"
)

type BaseController struct {
	beego.Controller

	IsLogin  bool
	UserInfo *models.User
}

type NestPreparer interface {
	NestPrepare()
}

type NestFinisher interface {
	NestFinish()
}

func (this *BaseController) SetLogin(user *models.User) {
	this.SetSession("userinfo", user.ID)
}

func (this *BaseController) DeleteLogin() {
	this.DelSession("userinfo")
}

func (this *BaseController) GetLogin() *models.User {
	user := &models.User{}
	user.ID = this.GetSession("userinfo").(uint)
	u, err := user.Find()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return &u
}

func (this *BaseController) Prepare() {
	fmt.Println("prepare")
	this.IsLogin = this.GetSession("userinfo") != nil
	fmt.Println(this.IsLogin)
	if this.IsLogin {
		this.UserInfo = this.GetLogin()
	}
	fmt.Println("xxxxx")

	this.Data["IsLogin"] = this.IsLogin
	this.Data["UserInfo"] = this.UserInfo

	if app, ok := this.AppController.(NestPreparer); ok {
		app.NestPrepare()
	}
}

func (this *BaseController) Finish() {
	if app, ok := this.AppController.(NestFinisher); ok {
		app.NestFinish()
	}
}
