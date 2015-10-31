package app

import (
	"github.com/astaxie/beego"

	"github.com/lei-cao/beefer/app/controllers"
)

func init() {
	beego.Router("/", &controllers.BeeferController{})
	beego.Router("/user/:user", &controllers.BeeferController{})
}
