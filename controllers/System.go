package controllers

import (
	"0934/models"
	"fmt"

	"github.com/astaxie/beego"
	//"strconv"
)

type SystemController struct {
	beego.Controller
}

//后台系统进入首页
func (this *SystemController) Index() {
	this.TplNames = "system/index.html"
	maps := models.GetSystemSetting()
	fmt.Println(maps)
}

//热更新系统
func (this *SystemController) UpdateSystem() {
	//models.UpdateSystemSetting()
}
