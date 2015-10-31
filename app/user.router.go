package app

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup;post:Signup")
	beego.Router("/user/logout", &UserController{}, "get:Logout")
}
