package models

import (
	//"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	//"log"
	//"reflect"
	//"errors"
)

func init() {
	orm.RegisterModel(new(SystemSetting), new(Page))
}

//系统设置信息
type SystemSetting struct {
	Id      int64
	Varname string
	Value   string
}

//y页面关键词
type Page struct {
	Id          int64
	Url_token   string //url
	Title       string //标题
	Keywords    string //关键词
	Description string //描述
	Content     string `orm:"null;type(text)"`
	Enabled     int8   //是否启用
}

//获取系统信息
func GetSystemSetting() []orm.Params {
	o := orm.NewOrm()
	ss := new(SystemSetting)
	var maps []orm.Params
	o.QueryTable(ss).Values(&maps)
	return maps
}

//更新系统信息
func UpdateSystemSetting(ss *SystemSetting) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(ss)
	return num, err
}

//添加关键词.....
func AddPage(page *Page) (int64, error) {
	o := orm.NewOrm()

	num, err := o.Insert(page)
	return num, err
}

//获取关键词
func GetPage(id int64) []orm.Params {
	o := orm.NewOrm()
	page := new(Page)
	var maps []orm.Params
	o.QueryTable(page).Filter("id", id).Values(&maps)
	return maps

}

//更新page
func UpdatePage(page *Page) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(page)
	return num, err
}
