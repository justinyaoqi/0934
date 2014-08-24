package controllers

import (
	"0934/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type QuestionController struct {
	beego.Controller
}

//首页
func (this *QuestionController) Index() {
	b := this.Ctx.GetCookie("username")
	be, _ := strconv.ParseBool(b)
	if be {
		username := this.Ctx.GetCookie("username")
		uid, _ := strconv.ParseInt(this.Ctx.GetCookie("uid"), 10, 64)
		user := models.GetUser(username, uid)
		this.Data["user"] = user
		list, _ := models.ListQuestion(1, 10, "question_id")
		this.Data["qlist"] = list
		this.TplNames = "question/index.html"
	} else {
		list, _ := models.ListQuestion(1, 10, "question_id")
		this.Data["qlist"] = list
		this.TplNames = "question/index.html"
		//this.TplNames = "question/index.html"
	}

}

//显示具体问题
func (this *QuestionController) Show() {
	maps := this.Ctx.Input.Params
	if maps["0"] == "lists" {
		qid, _ := strconv.ParseInt(maps["1"], 10, 64)
		qt := models.GetOneQustionByQid(qid)
		answer := models.GetAnswer(qid)
		qcomments := models.GetComments(qid)
		this.Data["answer"] = answer
		this.Data["qt"] = qt
		this.Data["qcomments"] = qcomments
		this.TplNames = "question/show.html"
	} else {
		this.Redirect("/", 302)
	}
}

//草稿
func (this *QuestionController) Draft() {
	BB := this.Ctx.GetCookie("username")

	if len(BB) > 0 {
		username := this.Ctx.GetCookie("username")
		uid, _ := strconv.ParseInt(this.Ctx.GetCookie("uid"), 10, 64)
		//用户信息模板
		user := models.GetUser(username, uid)
		this.Data["user"] = user
		//草稿模板
		dt := models.GetDraft(uid)
		this.Data["dt"] = dt
		this.TplNames = "question/draft.html"
	} else {
		this.Redirect("/question/draft", 301)
	}

}

//添加问答
func (this *QuestionController) AddQuestion() {
	Q := new(models.Question)

	qtitle := this.GetString("qtitle")

	qcontent := this.GetString("qcontent")
	//Q.Add_time = time.UTC

	Q.Question_content = qtitle
	Q.Question_detail = qcontent

	uid, err := strconv.ParseInt(this.GetString("uid"), 10, 64)
	//fmt.Println(Q.Add_time)
	if err != nil {
		log.Fatal(err)
	}
	Q.Published_uid = uid
	//fmt.Println(Q.Add_time)
	qid, err := models.AddQuestion(Q, uid)
	if err != nil && qid < 0 {
		log.Fatal(err)
	}
	this.Redirect("/question/index", 302)

}

//添加评论
func (this *QuestionController) AddComments() {
	qc := new(models.QuestionComments)
	published_uid, _ := strconv.ParseInt(this.GetString("published_uid"), 10, 64)
	qc.Uid = published_uid
	question_id, _ := strconv.ParseInt(this.GetString("question_id"), 10, 64)
	qc.Question_id = question_id
	qc.Add_time = time.Now()
	qccontent := this.GetString("qcomments")

	qc.Message = qccontent
	num, _ := models.AddComments(qc)
	if num > 0 {
		this.Redirect("/question/index", 302)
	}
}

//添加问题上传图片
func (this *QuestionController) Image() {
	_, header, err := this.GetFile("upfile")
	ext := strings.ToLower(header.Filename[strings.LastIndex(header.Filename, "."):])
	//fmt.Println(ext)
	out := make(map[string]string)
	out["url"] = ""
	out["fileType"] = ext
	out["original"] = header.Filename
	out["state"] = "SUCCESS"
	if err != nil {
		out["state"] = err.Error()
	} else {
		savepath := "./static/upload/" + time.Now().Format("20060102")
		//fmt.Println(savepath)
		if err := os.MkdirAll(savepath, os.ModePerm); err != nil {

			out["state"] = err.Error()
		} else {
			filename := fmt.Sprintf("%s/%d%s", savepath, time.Now().UnixNano(), ext)
			if err := this.SaveToFile("upfile", filename); err != nil {
				out["state"] = err.Error()
			} else {
				out["url"] = filename[1:]
			}
		}
	}

	editorid := this.GetString("editorid")
	this.Ctx.WriteString("<script>parent.UM.getEditor('" + editorid + "').getWidgetCallback('image')('" + out["url"] + "','" + out["state"] + "')</script>")

	this.Data["json"] = out
	this.ServeJson()

}

//添加赞功能
/*
func (this *QuestionController) AddZan() {
	//strid := this.GetString("uid")
	if this.Ctx.Input.IsAjax() {
		maps := this.Input()
		uid, _ := strconv.ParseInt(maps.Get("uid"), 10, 64)
		qid, _ := strconv.ParseInt(maps.Get("qid"), 10, 64)

	}

}
*/

