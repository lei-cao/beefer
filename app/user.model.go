package app

import (
	"github.com/astaxie/beego/orm"
)

func Init() {
	// register model
	orm.RegisterModel(new(User))
}

// The User model map to user table in db
type User struct {
	Id       int
	Username string `orm:"size(50)"`
	Email    string `orm:"size(200)"`
	Password string `orm:"size(100)"`
}
