package routers

import (
	"0934/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})   //根路由跳转
	beego.AutoRouter(&controllers.IndexController{})    //根
	beego.AutoRouter(&controllers.UserController{})     //用户设置目录跳转
	beego.AutoRouter(&controllers.QuestionController{}) //问答目录跳转
	beego.AutoRouter(&controllers.LoveController{})
	beego.AutoRouter(&controllers.AnswerController{})
	beego.AutoRouter(&controllers.SystemController{})
	beego.AutoRouter(&controllers.JobController{})

}
