package main

import (
	_ "github.com/georgehao/todo/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	logs.SetLogger("console")
	beego.Run()
}

