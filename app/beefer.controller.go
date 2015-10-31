package app

import "github.com/astaxie/beego"

type BeeferController struct {
	beego.Controller
	CurrentUser *User // Current logged in user
}

// The Prepare() controller method runs before the real action methods
func (c *BeeferController) Prepare() {
	s := c.GetSession(currentUserSessionKey)
	if s != nil {
		c.CurrentUser = s.(*User)
		c.Data["CurrentUser"] = c.CurrentUser
	}
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	name := c.GetString(":user")
	user := User{Username: name}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}
