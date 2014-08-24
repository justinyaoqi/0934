package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"reflect"
	//"errors"
	"time"
)

type Topic struct {
	Id                int64
	Topic_title       string
	Add_time          time.Time `orm:"null;type(datetime)"`
	Discuss_count     int64
	Topic_descreption string `orm:"null;type(text)"`
	Topic_pic         string
	Topic_lock        int8
	Forcus_count      int64 // 关注的计数
	User_related      int8  //是否被用户关联
	Url_token         string
	Merged_id         int64
	Seo_title         string
}

//话题关注数
type TopicFocus struct {
	Focus_id int64 `orm:"pk"`
	Topic_id int64
	Uid      int64
	Add_time time.Time `orm:"null;type(datetime)"`
}

//话题合并
type TopicMerge struct {
	Id        int64
	Source_id int64
	Target_id int64
	Uid       int64
	Add_time  time.Time `orm:"null;type(datetime)"`
}
type TopicRelation struct {
	Id       int64
	Topic_id int64
	Item_id  int64
	Add_time time.Time `orm:"null;type(datetime)"`
	Uid      int64
	Type     string
}

func init() {
	orm.RegisterModel(new(Topic), new(TopicFocus), new(TopicMerge), new(TopicRelation))
}

//增加话题
func AddTopic(t *Topic) (int64, error) {
	o := orm.NewOrm()
	//topic := new(Topic)
	num, err := o.Insert(t)
	return num, err
}

//删除话题
func DelTopic(id int64) (int64, error) {
	o := orm.NewOrm()
	t := new(Topic)
	num, err := o.Delete(t)
	return num, err
}

//更新话题
func UpdateTopic(t *Topic) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(t)
	return num, err
}

//获取用户关注的话题
func GetTopicFocus(uid int64) []orm.Params {
	o := orm.NewOrm()
	tf := new(TopicFocus)
	var maps []orm.Params
	o.QueryTable(tf).Filter("uid", uid).Values(&maps)
	return maps
}

//删除用户关注的话题
func DelTopicFocus(tfid int64) (int64, error) {
	o := orm.NewOrm()
	tf := new(TopicFocus)
	tf.Focus_id = tfid
	num, err := o.Delete(tf)
	return num, err
}
