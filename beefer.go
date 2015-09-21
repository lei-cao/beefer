package main

import "github.com/astaxie/beego"

const usernameCookieKey string = "username"

type BeeferController struct {
	beego.Controller
	CurrentUser *User // Current logged in user
}

// The Prepare() controller method runs before the real action methods
func (c *BeeferController) Prepare() {
	username := c.Ctx.GetCookie(usernameCookieKey)
	if username == "" {
		return
	}
	user := User{Username: username}
	c.CurrentUser = &user
	c.Data["CurrentUser"] = c.CurrentUser
}

type UserController struct {
	BeeferController
}

type User struct {
	Username string
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	name := c.GetString(":user")
	user := User{Username: name}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}

func (c *UserController) Login() {
	c.TplNames = "user/login.tpl"
}

func (c *UserController) Signup() {
	c.TplNames = "user/signup.tpl"
	if c.Ctx.Request.Method == "POST" {
		username := c.GetString("username")
		password := c.GetString("password")
		password2 := c.GetString("password2")
		if password != password2 {
			c.Data["ValidateMessage"] = "两次密码不一致"
			return
		}
		user := User{Username: username}
		c.Data["User"] = user
		c.Ctx.SetCookie(usernameCookieKey, username)
		c.Redirect("/", 302)
	}

}

func (c *UserController) Logout() {
	c.CurrentUser = nil
	c.Data["CurrentUser"] = c.CurrentUser
	c.Ctx.SetCookie(usernameCookieKey, "")
	c.TplNames = "beefer.tpl"
}

func main() {
	beego.Router("/", &BeeferController{})
	beego.Router("/user/:user", &BeeferController{})
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup;post:Signup")
	beego.Router("/user/logout", &UserController{}, "get:Logout")
	beego.Run(":8085")
}
