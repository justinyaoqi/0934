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

type School struct {
	School_id   int64 `orm:"pk"`
	School_type int16
	School_code int64
	School_name string
	Area_code   int64
}

func init() {
	orm.RegisterModel(new(School))
}

//获取学校信息
func GetSchool(id int64) []orm.Params {
	o := orm.NewOrm()
	s := new(School)
	var maps []orm.Params
	o.QueryTable(s).Filter("school_id", id).Values(&maps)
	return maps
}

//更新学校信息
func UpdateSchool(sl *School) (int64, error) {
	o := orm.NewOrm()
	//s := new(School)
	num, err := o.Update(sl)
	return num, err
}
