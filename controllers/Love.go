package controllers

import (
	//"0934/models"
	//"fmt"
	"github.com/astaxie/beego"
	//"strconv"
)

type LoveController struct {
	beego.Controller
}

//相亲首页
func (this *LoveController) Index() {
	this.TplNames = "love/index.html"
}

//同站登录
func (this *LoveController) Login() {
	this.TplNames = "love/login.html"
}

//退出
func (this *LoveController) Quit() {
	this.TplNames = "question/index.html"
}

//注册为相亲网站用户
func (this *LoveController) Regester() {
	this.TplNames = "love/reg.html"
}

//显示求亲用户
func (this *LoveController) Show() {
	this.TplNames = "love/show.html"
}

//相亲个人用户主页
func (this *LoveController) Home() {
	this.TplNames = "love/home.html"
}

//选择男或女
/*
func (this *LoveController) getListBySex() {

}
*/
