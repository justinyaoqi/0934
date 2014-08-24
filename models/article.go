package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"log"
	//"reflect"
	//"time"
)

type Articles struct {
	Id             int64
	Uid            int64  //关联的用户  id
	Title          string //用户标题
	Message        []byte //消息
	Comments       int64  //评论id
	Views          int64  //阅览人数
	Add_time       int64  //添加时间
	Has_attach     int8   //是否有附件
	Lock           int8   //是否锁定
	Votes          int64  //投票数
	Title_fulltext []byte //文章内容
	Categoty_id    int32  //分类id
}
type ArticleComment struct {
	Id         int64  //评论id
	Uid        int64  //关联用户id
	Article_id int64  //文章id
	Message    []byte //消息内容
	Add_time   int64  //添加时间
	At_uid     int64  //@的用户
	Votes      int64  //投票数
}
type ArticleVote struct {
	Id                int64  //文章赞成id
	Uid               int64  //关联的用户id
	Type              string //类型
	Item_id           int32  //分类id
	Rating            int8   //等级
	Time              int64  //投票时间
	Reputation_factor int32
	Item_uid          int64
}

//增加文章
func AddArticle(a *Articles) (id int64, err error) {
	o := orm.NewOrm()
	article := new(Articles)
	article.Title = a.Title
	article.Message = a.Message
	id, err = o.Insert(article)
	return id, err
}

//删除文章
func DelArticle(id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Articles{Id: id})
	return status, err

}
