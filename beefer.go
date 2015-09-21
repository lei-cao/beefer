package main

import "github.com/astaxie/beego"

const currentUserSessionKey string = "currentUser"

type BeeferController struct {
	beego.Controller
	CurrentUser *User // Current logged in user
}

// The Prepare() controller method runs before the real action methods
func (c *BeeferController) Prepare() {
	s :=c.GetSession(currentUserSessionKey)
	if s != nil {
		c.CurrentUser = s.(*User)
		c.Data["CurrentUser"] = c.CurrentUser
	}
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
		user := &User{Username: username}
		c.Data["User"] = user
		c.SetSession(currentUserSessionKey, user)
		c.Redirect("/", 302)
	}

}

func (c *UserController) Logout() {
	c.Data["CurrentUser"] = nil
	c.DelSession(currentUserSessionKey)
	c.TplNames = "beefer.tpl"
}

func main() {
	// Enable session. The default engine is memory
	beego.SessionOn = true
	beego.Router("/", &BeeferController{})
	beego.Router("/user/:user", &BeeferController{})
	beego.Router("/user/login", &UserController{}, "get:Login")
	beego.Router("/user/signup", &UserController{}, "get:Signup;post:Signup")
	beego.Router("/user/logout", &UserController{}, "get:Logout")
	beego.Run(":8085")
}
