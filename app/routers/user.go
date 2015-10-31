package app

import (
	"github.com/astaxie/beego"

	"github.com/lei-cao/beefer/app/controllers"
)

func init() {
	beego.Router("/user/login", &controllers.UserController{}, "get:Login")
	beego.Router("/user/signup", &controllers.UserController{}, "get:Signup;post:Signup")
	beego.Router("/user/logout", &controllers.UserController{}, "get:Logout")
}
