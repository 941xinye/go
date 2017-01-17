package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type Artivity struct {
	Id             int
	ActivityName   string
	ActivityUrl    string
	ActivityRemark string
	IsDelete       uint
	Inputtime      int
}

func (u *Artivity) TableName() string {
	return "pre_activity1"
}
func init() {
	orm.RegisterModel(new(Artivity))
}

func GetA(Id int) (obj Artivity) {
	o := orm.NewOrm()
	user := Artivity{Id: Id}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.ActivityName)
	}

	return user
}
