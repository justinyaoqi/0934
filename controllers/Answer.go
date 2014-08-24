package controllers

import (
	"0934/models"
	"fmt"
	"log"
	"github.com/astaxie/beego"
	//"os"
	//"strings"
	//"encoding/json"
	"strconv"
	"time"
)

type AnswerController struct {
	beego.Controller
}

//添加评论  Ajax方法回调
func (this *AnswerController) AddAnswer() {
	//fmt.Println(this.Input())
	maps := this.Input()
	answercommet := new(models.AnswerComments)
	answercommet.Answer_id, _ = strconv.ParseInt(maps.Get("Answer_id"), 10, 64)
	answercommet.Uid, _ = strconv.ParseInt(maps.Get("Uid"), 10, 64)
	answercommet.Message = maps.Get("Answer_content")
	answercommet.Time = time.Now()
	num, err := models.AddAnswerComments(answercommet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(num)
	this.Ctx.WriteString("1")
}
