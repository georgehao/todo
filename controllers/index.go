package controllers

import (
	"github.com/georgehao/todo/models"
	"strconv"
	"time"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	if !this.IsLogin {
		this.Ctx.Redirect(302, this.URLFor("LoginController.Login"))
		return
	}
	this.TplName = "index.html"

	var task models.Task
	task.UserId = this.UserInfo.ID
	per, _ := this.GetInt("per", models.PAGE_PER)
	pages, err := task.Page(per, this.SetPaginator(per, int64(len(task.All()))).Offset())
	if err != nil {
		panic(err)
	}

	this.Data["tasks"] = pages
	this.Data["total"] = len(task.All())
	this.Data["done"] = len(task.Done())
	this.Data["user"] = this.UserInfo.Email

	importants, err := task.MoreImportant()
	if err != nil {
		panic(err)
	}
	this.Data["importants"] = importants
}

func (this *IndexController) Add() {
	addContent := this.GetString("add-content")
	addDegree := this.GetString("add-degree")

	var task models.Task
	task.Content = addContent
	degree, err := strconv.Atoi(addDegree)
	if err != nil {
		panic(err)
	}
	task.Degree = degree
	task.CreatedAt = time.Now()
	task.UserId = this.UserInfo.ID

	_, err = task.Create()
	if err != nil {
		panic("Create Task error")
	}

	this.TplName = "index.html"
	this.ServeJSON()
}

func (this *IndexController) Modify() {
	var task models.Task
	id, err := strconv.Atoi(this.GetString("id"))
	if err != nil {
		panic(err)
	}
	task.ID = uint(id)

	task.Modify()
	this.ServeJSON()
}

func (this *IndexController) About() {
	if !this.IsLogin {
		this.Ctx.Redirect(302, this.URLFor("LoginController.Login"))
		return
	}
	this.TplName = "about.html"
	this.Data["user"] = this.UserInfo.Email
}