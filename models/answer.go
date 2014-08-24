package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"reflect"
	//	"errors"
	"time"
)

//数据库配置初始化
func init() {

	orm.RegisterModel(new(Answer), new(AnswerComments))
	orm.DefaultTimeLoc = time.UTC

}

type Answer struct {
	Answer_id          int64     `orm:"pk"`
	Question_id        int64     `orm:""`
	Answer_content     string    `orm:"null;type(text)"`
	Add_time           time.Time `orm:"null;type(datetime)"`
	Against_count      int64     `orm:"default(0)"`
	Agree_count        int64     `orm:"default(0)"`
	Uid                int64     `orm:"default(0)"`
	Comment_count      int64     `orm:"default(0)"`
	Uninterested_count int64     `orm:"default(0)"`
	Thanks_count       int64     `orm:"default(0)"`
	Category_id        int64     `orm:"default(0)"`
	Has_attach         int8      `orm:"default(0)"`
	Ip                 int64     `orm:"null"`
	Force_fold         int8      `orm:"default(0)"`
	Anonymous          int8      `orm:"default(0)"` //是否匿名回答
	Publish_source     string    `orm:"null;size(16)"`
}

type AnswerComments struct {
	Id        int64
	Answer_id int64
	Uid       int64
	Message   string    `orm:"null;type(text)"`
	Time      time.Time `orm:"null;type(datetime)"`
}
type AnswerThanks struct {
	Id        int64
	Uid       int64
	Answer_id int64
	//User_name string
	Time time.Time `orm:"null;type(datetime)"`
}

//不敢兴趣的问题
type AnswerUninsterrested struct {
	Id          int64     `orm:"pk"`
	Question_id int64     //话题id
	Uid         int64     //用户id
	Time        time.Time `orm:"null;type(datetime)"`
}

type AnswerVote struct {
	Voter_id          int64     `orm:"pk"`
	Answer_id         int64     //回复id
	Answer_uid        int64     //回复作者id
	Vote_uid          int64     //用户id
	Add_time          time.Time `orm:"null;type(datetime)"`
	Vote_value        int8      //反对,或支持  0,1
	Reputation_factor int64     //声望值
}

//增加回复 返回成功的id,如果err不为空
func AddAnswer(ar *Answer) (int64, error) {
	o := orm.NewOrm()
	//answer := new(Answer)
	aid, err := o.Insert(ar)
	return aid, err
}

//获取评论
func GetAnswer(aid int64) (as []orm.Params) {
	o := orm.NewOrm()
	answer := new(Answer)
	o.QueryTable(answer).Filter("answer_id", aid).Values(&as)
	return as
}

//评论回复
//成功返回id,
func AddAnswerComments(asc *AnswerComments) (int64, error) {
	o := orm.NewOrm()
	acid, err := o.Insert(asc)
	return acid, err
}

//获取评论回复

func GetAnswerComments(aid int64) (ac []orm.Params) {
	o := orm.NewOrm()
	answercomments := new(AnswerComments)
	o.QueryTable(answercomments).Filter("answer_id", aid).Values(&ac)
	return ac
}

//增加感谢评论

func AddAnswerThanks(aat *AnswerThanks) (int64, error) {
	o := orm.NewOrm()
	atid, err := o.Insert(aat)
	return atid, err
}

//增加不敢兴趣
func AddAnswerUninsterrested(aat *AnswerUninsterrested) (int64, error) {
	o := orm.NewOrm()
	atid, err := o.Insert(aat)
	return atid, err
}
