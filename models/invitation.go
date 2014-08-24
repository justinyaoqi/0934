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

func init() {
	orm.RegisterModel(new(Invitation))
}

//用户邀请表
type Invitation struct {
	Invitation_id    int64     `orm:"pk"` //邀请 id
	Uid              int64     //关联用户 id
	Invitation_code  string    //邀请码
	Invitation_email string    //邀请邮箱
	Add_time         time.Time //邀请时间
	Add_ip           int64     //邀请 ip
	Active_expire    int8      //激活超时
	Active_time      time.Time //激活时间
	Active_ip        int64     //激活 ip
	Active_status    int8      //激活状态
	Active_uid       int64     //新激活的用户  id
}

//添加邀请
func AddInvitation(invitation *Invitation) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Insert(invitation)
	return num, err
}

//获取邀请信息
func GetInvitation(uid int64) []orm.Params {
	o := orm.NewOrm()
	var maps []orm.Params
	invitation := new(Invitation)
	o.QueryTable(invitation).Filter("uid", uid).Values(&maps)
	return maps
}
