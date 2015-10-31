package controllers

import (
	"github.com/astaxie/beego"

	"github.com/lei-cao/beefer/app/models"
)

type BeeferController struct {
	beego.Controller
	CurrentUser *models.User // Current logged in user
}

// The Prepare() controller method runs before the real action methods
func (c *BeeferController) Prepare() {
	s := c.GetSession(currentUserSessionKey)
	if s != nil {
		c.CurrentUser = s.(*models.User)
		c.Data["CurrentUser"] = c.CurrentUser
	}
}

// The Get method to handle the GET request
func (c *BeeferController) Get() {
	name := c.GetString(":user")
	user := models.User{Username: name}
	c.Data["User"] = user

	c.TplNames = "beefer.tpl"
}
