package controllers

import (
	"github.com/astaxie/beego/orm"

	"github.com/lei-cao/beefer/app/models"
)

const currentUserSessionKey string = "currentUser"

type UserController struct {
	BeeferController
}

// The Get method to handle the GET request
func (c *UserController) Get() {
	name := c.GetString(":user")
	user := models.User{Username: name}
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
		user := &models.User{Username: username, Password: password}

		o := orm.NewOrm()
		if created, id, err := o.ReadOrCreate(user, "Username"); err == nil {
			if created {
				user.Id = int(id)
			} else {
				c.Data["ValidateMessage"] = "User existed."
				return
			}
		} else {
			c.Data["ValidateMessage"] = err
			return
		}

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
