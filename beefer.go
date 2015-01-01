package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

type User struct {
	Username string
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	user := User{Username: "Alice"}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}

func (c *UserController) Login() {
	c.TplNames = "beefer.tpl"
}

func (c *UserController) Signup() {
	c.TplNames = "beefer.tpl"
}

func (c *UserController) Logout() {
	c.TplNames = "beefer.tpl"
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup")
	beego.Router("/user/logout", &UserController{}, "post:Logout")
	beego.Run(":8088")
}
