package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"log"
	//"reflect"
	"time"
	//"errors"
	//"strconv"
	//"io"
	//"github.com/beego/social-auth"
	//"github.com/beego/social-auth/apps"
)

//数据库配置初始化
func init() {

	orm.RegisterModel(new(Users), new(ActiveData), new(UserActionHistoryData), new(UserActionHistory), new(UsersActionHistoryFresh), new(UsersOnline), new(UsersSina), new(UserFollow), new(UsersAttrib), new(UsersGroup), new(UsersNotificationSetting))
	//orm.RegisterDataBase("default", "mysql", "root:dskdsk@/0934?charset=utf8")
}

//用户类
type Users struct {
	Uid                    int64     `orm:"pk"`
	User_name              string    `orm:"unique;column(user_name);null;size(255)"`
	Email                  string    `orm:"null;size(255)"`
	Mobile                 string    `orm:"null;size(16)"`
	Password               string    `orm:"null;size(32)"`
	Salt                   string    `orm:"null;size(16)"`
	Avatar_file            string    `orm:"column(avatar_file);null;size(127)"`
	Sex                    int8      `orm:"null;"`
	Birthday               int32     `orm:"null"`
	Province               string    `orm:"null"`
	City                   string    `orm:"null"`
	Job_id                 string    `orm:"column(job_id);default(0)"`
	Reg_time               time.Time `orm:"null;type(datetime)"`
	Reg_ip                 int64     `orm:"column(reg_ip);null"`
	Last_login             int32     `orm:"column(last_login);`
	Last_ip                int64     `orm:"column(last_ip);null"`
	Online_time            int32     `orm:column(online_time);"default(0)"`
	Last_active            int32     `orm:"column(last_active);null"`
	Notification_unread    int32     `orm:"column(notification_unread);default(0)"`
	Inbox_unread           int32     `orm:"column(inbox_unread);default(0)"`
	Inbox_recv             int8      `orm:"column(inbox_recv);default(0)"`
	Fans_count             int32     `orm:"column(fans_count);default(0)"`
	Friend_count           int32     `orm:"column(friend_count);default(0)"`
	Invite_count           int32     `orm:"column(invite_count);default(0)"`
	Question_count         int32     `orm:"default(0)"`
	Answer_count           int32     `orm:"default(0)"`
	Topic_focus_count      int32     `orm:"default(0)"`
	Invitation_available   int32     `orm:"default(0)"`
	Group_id               int16     `orm:"default(0)"`
	Reputation_group       int16     `orm:"default(0)"`
	Forbidden              int8      `orm:"default(0)"`
	Valid_email            int8      `orm:"default(0)"`
	Is_first_login         int8      `orm:"default(1)"`
	Agree_count            int32     `orm:"default(0)"`
	Thanks_count           int32     `orm:"default(0)"`
	Views_count            int32     `orm:"default(0)"`
	Reputation_update_time time.Time `orm:"null;type(datetime)"`
	Weibo_visit            int8      `orm:"default(1)"`
	//Integral               int32  `orm:"default(0)"`
	Draft_count      int64  `orm:"default(0)"`
	Common_email     string `orm:"null;size(255)"`
	Url_token        string `orm:"null;size(32)"`
	Url_token_update int32  `orm:"default(0)"`
	Verified         string `orm:"null;size(32)"`
	Default_timezone string `orm:"null";size(32)`
	Email_settings   string `orm:"null;size(255)"`
	Weixin_settings  string `orm:"null;size(255)"`
	Recent_topics    string `orm:"null;type(text)"`
	Islover          int8   `orm:"null"`
}

//用户登录

func IsLogin(uname string, upass string) (int64, int) {

	o := orm.NewOrm()
	var users Users
	users.User_name = uname
	users.Password = upass
	uid, err := o.QueryTable("users").Filter("user_name", users.User_name).Filter("password", users.Password).All(&users, "Uid")
	if err == orm.ErrMultiRows {
		return uid, 2
	}
	if err == orm.ErrNoRows {
		return uid, 0
	}
	return uid, 1
}

//根据用户id获取一个用户
func GetOne(uid int64) (maps []orm.Params) {
	o := orm.NewOrm()
	user := Users{}
	o.QueryTable(user).Filter("uid", uid).Values(&maps)
	//fmt.Println(maps)
	//fmt.Println(user.Userfollows.User)
	return maps

}

//获取用户分页
func GetUserList(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(Users)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//增加用户
func AddUser(u *Users) (id int64, err error) {
	o := orm.NewOrm()
	user := new(Users)
	user = u
	id, err = o.Insert(user)
	return id, err

}

//用户更新
func UpdateUser(u *Users) (num int64, err error) {
	o := orm.NewOrm()
	user := make(orm.Params)
	//user.Uid = u.Uid
	num, err = o.QueryTable("users").Filter("uid", u.Uid).Update(user)
	return num, err
}

//删除用户
func DelUserById(uid int64) (status int64, err error) {
	o := orm.NewOrm()
	return o.Delete(&Users{Uid: uid})
}

//根据用户名获取用户信息
func GetUser(user_name string, uid int64) (users []orm.Params) {
	user := Users{} //用户信息
	userfollow := UserFollow{}
	o := orm.NewOrm()
	o.QueryTable(user).Filter("user_name", user_name).Limit(1).Values(&users)
	fansnum, _ := o.QueryTable(userfollow).Filter("fans_uid", uid).Count()  //我的关注
	frinum, _ := o.QueryTable(userfollow).Filter("friend_uid", uid).Count() //我的粉丝
	users[0]["Fans_count"] = fansnum
	users[0]["Friend_count"] = frinum

	return users
}

//用户属性
type UsersAttrib struct {
	Id           int64  `orm:"pk;_;"`
	Uid          int64  `orm:"null"`
	Introduction string `orm:"null"`
	Signature    string `orm:"null"`
	Qq           string `orm:"null"`
	Homepage     string `orm:"null"`
}

func (ua *UsersAttrib) TableName() string {
	return "users_attrib"
}

//获取用户属性
func GetUserAttrib(uid int64) []orm.Params {
	o := orm.NewOrm()
	ua := new(UsersAttrib)
	var maps []orm.Params
	o.QueryTable(ua).Filter("uid", uid).Values(&maps)
	return maps

}

//更改用户属性
func UpdateUserAttrib(ua *UsersAttrib) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(ua)
	return num, err
}

//删除用户属性
func DelUserAttrib(uid int64) (int64, error) {
	o := orm.NewOrm()
	ua := new(UsersAttrib)
	ua.Uid = uid
	num, err := o.Delete(ua)
	return num, err
}

//用户关注数
type UserFollow struct {
	Follow_id  int64 `orm:"pk"`
	Fans_uid   int64
	Friend_uid int64
	Add_time   time.Time `orm:"null;type(datetime)"`
}

func (uf *UserFollow) TableName() string {
	return "user_follow"
}

//粉丝添加

//用户组
type UsersGroup struct {
	Group_id          int32   `orm:"pk"`              //组id
	Type              int16   `orm:"default(0)"`      //类型
	Custom            int16   `orm:"default(0)"`      //习惯
	Group_name        string  `orm:"null"`            //组名
	Reputation_lower  int32   `orm:"default(0)"`      //声望级别
	Reputation_higer  int32   `orm:"default(0)"`      //声望高度
	Reputation_factor float32 `orm:"default(0)"`      //声望因素
	Permission        string  `orm:"null;type(text)"` //权限
}

func (ug *UsersGroup) TableName() string {
	return "users_group"
}

//短消息通知
type UsersNotificationSetting struct {
	Notice_setting_id int64  `orm:"pk"`
	Uid               int64  `orm:"null"` //用户id
	Data              string `orm:"null"` //数据内容
}

func (us *UsersNotificationSetting) TableName() string {
	return "users_notification_setting"
}

//用户行踪
type UserActionHistory struct {
	History_id         int64     `orm:"pk"`
	Uid                int64     //用户id
	Associate_type     int8      //关联类型: 1 问题 2 回答 3 评论 4 话题   Associate关联
	Associate_action   int16     //操作类型
	Associate_id       int64     //关联id
	Add_time           time.Time `orm:"null;type(datetime)"` //添加时间
	Associate_attached int64     //附件
	Anonymous          int8      //是否匿名
	Fold_status        int8      //折叠状态
}

func (uac *UserActionHistory) TableName() string {
	return "user_action_history"
}

//用户行为习惯
type UserActionHistoryData struct {
	History_id         int64  `orm:"pk"`              //历史id
	Associate_content  string `orm:"null;type(text)"` //关系内容
	Associate_attached string `orm:"null;type(text)"` //关系附件
	Addon_data         string `orm:"null;type(text)"` //关系数据
}

func (uahd *UserActionHistoryData) TableName() string {
	return "user_action_history_data"
}

//用户流
type UsersActionHistoryFresh struct {
	Id             int64
	History_id     int64     //用户行为历史id
	Associate_id   int64     //关系id
	Associate_type int8      //关系类型
	Add_time       time.Time `orm:"null;type(datetime)"` //增加时间
	Uid            int64     //用户id
	Anonymous      int8      //是否匿名
}

func (uahf *UsersActionHistoryFresh) TableName() string {
	return "user_action_history_fresh"
}

//在线用户统计
type UsersOnline struct {
	Id          int64 `orm:"pk"`
	Uid         int64
	Last_active int    `orm:"default"`
	Ip          int64  `orm:"null"`
	Active_url  string `orm:"null;size(255)"`
	User_agent  string `orm:"null:size(255)"`
}

func (uo *UsersOnline) TableName() string {

	return "users_online"
}

//用户QQ绑定设置
type UsersQq struct {
	Id                 int64     `orm:"pk"`
	Uid                int64     `orm:"null"`
	Type               string    `orm:"null:size(20)"`
	Name               string    `orm:"null:size(64)"`
	Location           string    `orm:"null:size(255)"`
	Gender             string    `orm:"null"`
	Add_time           time.Time `orm:"null;type(datetime)"`
	access_token       string    `orm:"null:size(64)"`
	Oauth_token_secret string    `orm:"null:size(64)"`
	Nick               string    `orm:"null:size(64)"`
}

func (uq *UsersQq) TableName() string {
	return "users_qq"
}

//QQ登录
func QQLogin() {}

//新浪微博设置
type UsersSina struct {
	Id                int64
	Uid               int64
	Name              string    `orm:"null:size(64)"`
	Location          string    `orm:"null:size(255)"`
	Description       string    `orm:"null;type(text)"`
	Url               string    `orm:"null:size(255)"`
	Profile_image_url string    `orm:"null:size(255)"`
	Gender            string    `orm:"null:size(8)"`
	Add_time          time.Time `orm:"null;type(datetime)"`
	Access_token      string    `orm:"null:size(64)"`
}

func (us *UsersSina) TableName() string {
	return "users_sina"
}

//微信设置

type UserWeixin struct {
	Id              int64
	Uid             int64
	Openid          string
	Expires_in      int64
	Access_token    string
	Refresh_taken   string
	Scope           string
	Headimgurl      string
	Nickname        string
	sex             int
	Province        string
	City            string
	Country         string
	Add_time        time.Time `orm:"null;type(datetime)"`
	Latitude        float32
	Longtitude      float64
	Location_update int32
}

//用户激活
type ActiveData struct {
	Active_id        int64 `orm:"pk"`
	Uid              int64
	Expire_time      time.Time `orm:"null;type(datetime)"`
	Add_time         time.Time `orm:"null;type(datetime)"`
	Active_code      string
	Active_type_code string
	Add_ip           int64
	Active_time      time.Time `orm:"null;type(datetime)"`
	Active_ip        int64
}

func (ad *ActiveData) TableName() string {
	return "active_data"
}

//func GetUser(uname string, uid int64) u *users{}
type Jobs struct {
	Id       int64
	Job_name string
}

/*
func Test(uid int64) (u []orm.Params) {
	user := new(Users)
	o := orm.NewOrm()
	o.QueryTable(user).Filter("uid", uid).RelatedSel().Values(&u)
	//fmt.Println(u)
	return u
}
*/
