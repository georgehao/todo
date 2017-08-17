package controllers

import (
	"fmt"
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
	per, _ := this.GetInt("per", models.PAGE_PER)
	pages, err := task.Page(per, this.SetPaginator(per, int64(len(task.All()))).Offset())
	if err != nil {
		panic(err)
	}

	this.Data["tasks"] = pages
	this.Data["total"] = len(task.All())
	this.Data["done"] = len(task.Done())

	importants, err := task.MoreImportant()
	if err != nil {
		panic(err)
	}
	this.Data["importants"] = importants
}

func (this *IndexController) Add() {
	fmt.Println("index controller add")
	addContent := this.GetString("add-content")
	addDegree := this.GetString("add-degree")

	fmt.Println("content", addContent)
	fmt.Println("degree", addDegree)

	var task models.Task
	task.Content = addContent
	degree, err := strconv.Atoi(addDegree)
	if err != nil {
		panic(err)
	}
	task.Degree = degree
	task.CreatedAt = time.Now()

	_, err = task.Create()
	if err != nil {
		panic("Create Task error")
	}

	this.TplName = "index.html"
	this.ServeJSON()
}

func (this *IndexController) Modify() {
	fmt.Println("modify")
	var task models.Task
	id, err := strconv.Atoi(this.GetString("id"))
	if err != nil {
		panic(err)
	}
	task.ID = uint(id)

	task.Modify()
	this.ServeJSON()
}