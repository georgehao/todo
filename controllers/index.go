package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}

func (this *IndexController) Index() {
	fmt.Println("index controller get")
	this.TplName = "index.html"
	this.Render()
}