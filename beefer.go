package main

import (
	"github.com/astaxie/beego"

	_ "github.com/lei-cao/beefer/app"
)

func main() {
	// Enable session. The default engine is memory
	beego.SessionOn = true
	beego.Run(":8085")
}
