package main

import "github.com/astaxie/beego"

type BeeferController struct {
	beego.Controller
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	c.Ctx.WriteString("Hello Beego")
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":8085")
}
