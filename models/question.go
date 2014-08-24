package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"reflect"
	"errors"
	"time"
)

type Question struct {
	Question_id               int64     `orm:"pk"`
	Question_content          string    `orm:"null;size(255)"`
	Question_detail           string    `orm:type(text)`
	Add_time                  time.Time `orm:"type(datetime);index"`
	Update_time               time.Time `orm:"null;type(datetime)"`
	Published_uid             int64     `orm:"null"`
	Answer_count              int64     `orm:"default(0)"`
	Answer_users              int64     `orm:"default(0)"`
	View_count                int64     `orm:"default(0)"`
	Focus_count               int64     `orm:"default(0)"`
	Comment_count             int64     `orm:"default(0)"`
	Action_history_id         int64     `orm:"default(0)"`
	Category_id               int64     `orm:"default(0)"`
	Agree_count               int64     `orm:"default(0)"`
	Against_count             int64     `orm:"default(0)"`
	Best_answer               int64     `orm:"default(0)"`
	Has_attach                int8      `orm:"default(0)"`
	Unverified_modify         string    `orm:"null;type(text)"`
	Ip                        int64     `orm:"null"`
	Last_answer               int64     `orm:"default(0)"`
	Popular_value             float64   `orm:"default(0)"`
	Popular_value_update      float64   `orm:"default(0)"`
	Lock                      int8      `orm:"default(0)"`
	Anonymous                 int8      `orm:"default(0)"`
	Thanks_count              int64     `orm:"default(0)"`
	Question_content_fulltext string    `orm:"null;type(text)"`
	Is_recommend              int8      `orm:"default(0)"`
}

type QuestionThanks struct {
	Id          int64
	Uid         int64     `orm:"default(0)"`     //用户id
	Question_id int64     `orm:"default(0)"`     //问题id
	User_name   string    `orm:"null;size(255)"` //想感谢的用户名
	Add_time    time.Time `orm:"null;type(datetime)"`
}

//回复

// 赞成
type Approval struct {
	Id   int64
	Type string
	Data string `orm:"null;type(text)"`
	Uid  int64
	Time int64
}

//附件
type Attach struct {
	Id            int64
	File_name     int64
	Access_key    int64
	Add_time      int64
	File_location string
	Is_image      int8
	Item_type     string
	Item_id       int64
	Wait_approval int8
}

//分类表

type Category struct {
	Id        int64
	Title     string
	Type      string
	Icon      string //图标位置
	Parent_id int64  //父级分类id
	Sort      int16
	Url_token string
}

//数据库配置初始化
func init() {

	orm.RegisterModel(new(Question), new(QuestionComments), new(QuestionFocus), new(QuestionInVite), new(QuestionThanks), new(QuestionUninterested), new(Draft))

}
func (Q *Question) TableName() string {
	return "question"
}

//获取一个question
func GetOneQustionByQid(qid int64) (queston []orm.Params) {
	o := orm.NewOrm()
	Q := new(Question)
	o.QueryTable(Q).Filter("question_id", qid).Limit(1).Values(&queston)
	return queston
}

//增加问答
func AddQuestion(q *Question, published_id int64) (int64, error) {
	o := orm.NewOrm()
	question := new(Question)
	//var question Questions
	question.Question_content = q.Question_content
	question.Question_detail = q.Question_detail
	question.Published_uid = q.Published_uid
	time.LoadLocation("Asia/Beijing")

	question.Add_time = time.Now()
	id, err := o.Insert(question)
	return id, err
}

//问答查询,分页处理
func ListQuestion(page int64, page_size int64, sort string) (questions []orm.Params, count int64) {
	o := orm.NewOrm()
	question := new(Question)
	qs := o.QueryTable(question)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&questions)
	count, _ = qs.Count()
	return questions, count
}

//删除问答
func DelQuestion(qid int64) (int64, error) {
	o := orm.NewOrm()
	//question := new(Question)
	num, err := o.Delete(&Question{Question_id: qid})
	return num, err
}

//更新问答主表
func UpdateQuestion(qid int64) (int64, error) {
	o := orm.NewOrm()
	question := new(Question)
	//question.Question_id=qid
	if o.Read(question) == nil {
		question.Question_id = qid
		if num, err := o.Update(question); err != nil {

			//errors.New("更新失败")
			return num, err
		}
	}
	return 0, errors.New("更新成功")
}

//问题评论
type QuestionComments struct {
	Id          int64
	Question_id int64     // 问题id
	Uid         int64     //用户回复id
	Message     string    `orm:type(text)`
	Reply_count int64     //回复统计
	Add_time    time.Time `orm:"type(datetime);index"` //回复时间
}

//添加评论功能
func AddComments(c *QuestionComments) (int64, error) {
	//com := new(QuestionComments)
	o := orm.NewOrm()
	//o.QueryTable(com).Filter(c).Values(&comments)
	num, err := o.Insert(c)
	return num, err
}

//获取评论
func GetComments(qid int64) (qc []orm.Params) {
	o := orm.NewOrm()
	qcomments := new(QuestionComments)
	//qcomments.Question_id = qid
	o.QueryTable(qcomments).Filter("question_id", qid).Values(&qc)
	return qc

}

//添加赞功能
func AddZan(qid, uid int64) (int64, error) {
	o := orm.NewOrm()
	q := new(Question)
	q.Agree_count = q.Agree_count + 1
	q.Question_id = qid
	q.Published_uid = uid
	//o.QueryTable(q).Filter("question_id", qid).
	num, err := o.Update(q)
	return num, err
}

//不敢兴趣,没有帮助
type QuestionUninterested struct {
	Interested_id            int64 `orm:"pk"`
	Question_id              int64
	Uid                      int64
	Question_uninterestedcol string    `orm:"null;size(45)"`
	Add_time                 time.Time `orm:"null;type(datetime)"`
}

//没有帮助
func AddNoHelp(qid, uid int64) (int64, error) {
	o := orm.NewOrm()
	qu := new(QuestionUninterested)
	qu.Question_id = qid
	qu.Uid = uid
	qu.Add_time = time.Now()
	num, err := o.Insert(qu)
	return num, err
}

type QuestionFocus struct {
	Focus_id    int64 `orm:"pk"`
	Question_id int64
	Uid         int64
	Add_time    time.Time `orm:"null;type(datetime)"`
}

//收藏
func AddQuestionFocus(qid, uid int64) (int64, error) {
	o := orm.NewOrm()
	qf := new(QuestionFocus)
	qf.Uid = uid
	qf.Question_id = qid
	qf.Add_time = time.Now()
	num, err := o.Insert(qf)
	return num, err
}

//添加举报
//添加感谢
func AddQuestionThanks(qid, uid int64) (int64, error) {
	o := orm.NewOrm()
	qt := new(QuestionThanks)
	qt.Question_id = qid
	qt.Uid = uid
	qt.Add_time = time.Now()
	num, err := o.Insert(qt)
	return num, err
}

type QuestionInVite struct {
	Question_invite_id int64     `orm:"pk"`
	Question_id        int64     //问题id
	Sender_id          int64     //发送用户
	Recipients_uid     int64     //接受id
	Email              string    //邮箱
	Add_time           time.Time `orm:"null;type(datetime)"`
	Available_time     time.Time `orm:"null;type(datetime)"`
}

//添加邀请功能
func AddQuestionInvite(qi *QuestionInVite) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Insert(qi)
	return num, err
}

//草稿
type Draft struct {
	Id       int64
	Uid      int64
	Type     string `orm:"null";size(16)`
	Item_id  int64
	Data     string    `orm:"null;type(text)"`
	Add_time time.Time `orm:"null;type(datetime)"`
}

func GetDraft(uid int64) (dt []orm.Params) {
	o := orm.NewOrm()
	o.QueryTable("draft").Filter("uid", uid).Values(&dt)

	return dt

}
