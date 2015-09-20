package main

import (
	"text/template"

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
	var tpl string = `
        <html>
            <head>
                <title>Beefer!</title>
            </head>
            <body>
                <strong>Hello, {{.User.Username}}</strong>
            </body>
        </html>
    `
	data := make(map[interface{}]interface{})
	user := User{Username: "Alice"}
	data["User"] = user

	t := template.New("Beefer Template")
	t = template.Must(t.Parse(tpl))

	t.Execute(c.Ctx.ResponseWriter, data)
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Run(":8085")
}
