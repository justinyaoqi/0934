package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"log"
	//"reflect"
	//"time"
	//"errors"
	//"strconv"
	//"io"
)

//声望和级别
//用户声望

type ReputationCategory struct {
	Id             int64
	Uid            int64
	Category_id    int16 //级别id
	Update_time    int64 //更新时间
	Reputation     int64 //声望
	Thanks_count   int64 //感谢统计
	Agree_count    int64 //同意统计
	Question_count int64 //问题统计
}
type ReputationTopic struct {
	Id           int64
	Uid          int64 //用户id
	Topic_id     int64 //话题id
	Topic_count  int64 //话题统计
	Update_time  int64 //更新时间
	Agree_count  int64 //同意统计
	Thanks_count int64 //感谢统计
	Reputation   int64 //声望
}

func init() {
	orm.RegisterModel(new(ReputationCategory), new(ReputationTopic))
}

//获取

func GetReputationCategory(uid int64) []orm.Params {
	o := orm.NewOrm()
	var rc ReputationCategory
	var maps []orm.Params
	o.QueryTable(rc).Filter("uid", uid).Values(&maps)
	return maps
}

//获取声望话题
func GetReputationTopic(uid int64) []orm.Params {

	o := orm.NewOrm()
	var rt ReputationTopic
	var maps []orm.Params
	o.QueryTable(rt).Filter("uid", uid).Values(&maps)
	return maps
}
