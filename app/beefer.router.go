package app

import (
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &BeeferController{})
	beego.Router("/user/:user", &BeeferController{})
}
