package controllers

import (
	//"0934/models"
	//"fmt"
	"github.com/astaxie/beego"
)

type IndexController struct {
	beego.Controller
}
type UserPass struct {
	Name     string
	Password string
}

//首页跳转
func (this *IndexController) Get() {
	this.TplNames = "index.html"
}

//同为首页跳转
func (this *IndexController) Index() {

	this.TplNames = "index.html"
}
