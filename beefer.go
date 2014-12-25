package main

import (
	"github.com/astaxie/beego"
)

type BeeferController struct {
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

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":8080")
}
