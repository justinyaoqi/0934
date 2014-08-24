package main

import (
	_ "0934/routers"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {

	//orm.Debug = true
	orm.RegisterDataBase("default", "mysql", "root:dskdsk@/0934?charset=utf8")
	orm.DefaultTimeLoc = time.UTC
	beego.SessionOn = true
	beego.Run()
}
