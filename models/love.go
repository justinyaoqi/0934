package models

import (
	"fmt"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	//"log"
	//"reflect"
	//"time"
	//"errors"
	//"strconv"
	//"io"
)

type Lover struct {
	Id        int64
	Uid       int64
	Realname  string //真名
	Sex       int8   //性别
	Day       string //出生日期
	Hight     string //身高
	Education string //教育程度
	Where     string //出生地
	Income    string //收入
	Idcard    int64  //身份证 id
	Phone     string //用户手机
	Nickname  string //昵称
}

type Loverblog struct {
	id       int64  //
	Title    string //抢亲页展示标题,以引起其他人注意
	Blog     string `orm:"null;type(text)"`
	Url      string //用户url地址
	Loveruid int64  //关联用户 id
}

/*
func (l *Lover) TableName() string {
	return "lover"


}
func (Lb *Loverblog) TableName() string {
	return "loverblog"
}
func init() {
	orm.RegisterModel(new(Lover), new(Loverblog))
}
*/
//注册用户
/*
func Regester(lover *Lovers) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(Lover)
	return id, error

}
*/
//用户登录
//传入全局用户uid
//func LoverIsLogin(uid) {
//	IsLogin(uid)
//}
func SyscUser() {
	fmt.Println("hello")
}
