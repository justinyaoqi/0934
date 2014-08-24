package controllers

import (
	"0934/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {

	this.TplNames = "user/index.html"
}
func (this *UserController) Login() {
	uname := this.Input().Get("uname")
	upass := this.Input().Get("upass")
	uid, err := models.IsLogin(uname, upass)
	this.Ctx.SetCookie("uid", strconv.FormatInt(uid, 10), 36000, "/")
	this.Ctx.SetCookie("username", uname, 36000, "/")
	if err == 1 {
		this.Redirect("index", 302)
	} else {
		fmt.Println("登录失败")
		//}
		this.Redirect("index", 302)
		this.TplNames = "/"
	}

}

//获取用户消息
func (this *UserController) Message() {

	this.TplNames = "user/message.html"
}

//用户注册路由
func (this *UserController) Regester() {

	//fmt.Println(models.IsLogin("yaoqi", "dskdsk"))

	this.TplNames = "user/regester.html"
}

//跳转用户主页
func (this *UserController) Home() {
	this.TplNames = "user/home.html"
}

//修改用户密码
/*
func (this *UserController) changePass() {

	newpass := this.Input().Get("")

}
*/
//以下用户资料修改
