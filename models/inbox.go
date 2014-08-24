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

//私信
type Inbox struct {
	Id               int64
	Uid              int64  //发送者id
	Dialog_id        int64  //接受者id
	Message          string `orm:"null;type(text)"`
	Add_time         time.Time
	Sender_remove    int8 //发送者移除 是否
	Recipient_remove int8 //接受者移除  是否
	Receipt          int64
}

//聊天窗口
type InboxDialog struct { //消息对话
	Id               int64
	Sender_uid       int64     //发送者id
	Sender_unread    int64     //发送未读
	Recipient_uid    int64     //接受者id
	Recipient_unread int64     //接收者未读
	Add_time         time.Time //添加时间
	Update_time      time.Time //更新时间
	Sender_count     int64     //发送统计
	Recipient_count  int64     //接收统计
}

func init() {
	orm.RegisterModel(new(Inbox), new(InboxDialog))
}

//添加消息
func InsertInbox(i *Inbox) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Insert(i)
	return num, err
}

//获取消息
func GetInbox(uid int64) []orm.Params {
	o := orm.NewOrm()
	inbox := new(Inbox)
	var maps []orm.Params
	o.QueryTable(inbox).Filter("uid", uid).Values(&maps)
	return maps
}

//添加聊天
func AddInboxDialog(inboxdialog *InboxDialog) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Insert(inboxdialog)
	return num, err

}

//获取聊天信息by Sender_uid
func GetInboxDialogBySender(uid int64) []orm.Params {
	o := orm.NewOrm()
	inboxdialog := new(InboxDialog)
	var maps []orm.Params
	o.QueryTable(inboxdialog).Filter("sender_uid", uid).Values(&maps)
	return maps
}

//获取聊天信息by Recipient_uid
func GetInboxDialogByRecipient(uid int64) []orm.Params {
	o := orm.NewOrm()
	inboxdialog := new(InboxDialog)
	var maps []orm.Params
	o.QueryTable(inboxdialog).Filter("recipient_uid", uid).Values(&maps)
	return maps
}
