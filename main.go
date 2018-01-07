package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/toolbox"
	_ "github.com/mattn/go-sqlite3"
	_ "beego-ripple/routers"
	_ "beego-ripple/tasks"
)


func init() {
	fmt.Println("Init Main")
	orm.Debug = true

	orm.RegisterDriver("sqlite3", orm.DRSqlite)
	orm.RegisterDataBase("default", "sqlite3", "file:database/ripple.db")

	orm.RunCommand()
}

func main() {

	fmt.Println("Start Main")



	//user := User{Name: "slene"}
	//
	//// insert
	//id, err := o.Insert(&user)
	//fmt.Printf("ID: %d, ERR: %v\n", id, err)
	//
	//// update
	//user.Name = "astaxie"
	//num, err := o.Update(&user)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)
	//
	//// read one
	//u := User{Id: user.Id}
	//err = o.Read(&u)
	//fmt.Printf("ERR: %v\n", err)
	//
	//// delete
	//num, err = o.Delete(&u)
	//fmt.Printf("NUM: %d, ERR: %v\n", num, err)





	beego.Run()

	defer toolbox.StopTask()
}

