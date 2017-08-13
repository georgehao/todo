package controllers

import (
	"fmt"
	"github.com/georgehao/todo/models"
	"time"
	"strconv"
)

type IndexController struct {
	BaseController
}

func (this *IndexController) Index() {
	fmt.Println("index controller get")
	this.TplName = "index.html"

	var task models.Task
	this.Data["tasks"] = task.All()
	this.Data["total"] = len(task.All())
	this.Data["done"] = len(task.Done())
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
