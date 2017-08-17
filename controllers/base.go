package controllers

import (
	"github.com/astaxie/beego"
	"github.com/georgehao/todo/models"
	"github.com/georgehao/todo/lib"
	"fmt"
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

func (this *BaseController) SetPaginator(per int, nums int64) *lib.Paginator {
	p := lib.NewPaginator(this.Ctx.Request, per, nums)
	this.Data["paginator"] = p
	return p
}

func (this *BaseController) Prepare() {
	this.IsLogin = this.GetSession("userinfo") != nil
	if this.IsLogin {
		this.UserInfo = this.GetLogin()
	}

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
