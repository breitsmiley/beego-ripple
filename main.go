package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/mattn/go-sqlite3"
	_ "beego-ripple/routers"
	_ "beego-ripple/tasks"
)


func init() {

	beego.SetLogger("file", `{"filename":"logs/app.log"}`)
	beego.SetLogFuncCall(true)


	orm.Debug = true
	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:database/ripple.db")

	orm.RunCommand()
}

func main() {

	beego.Info("Start App")

	beego.Run()


	defer toolbox.StopTask()
}

