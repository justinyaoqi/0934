package controllers

import (
	//"0934/models"
	//"fmt"
	"github.com/astaxie/beego"
	//"strconv"
)

type JobController struct {
	beego.Controller
}

//求职信息展览
func (this *JobController) Index() {
	this.TplNames = "job/index.html"
}
