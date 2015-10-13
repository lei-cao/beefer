package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
)

const currentUserSessionKey string = "currentUser"

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

type UserController struct {
	BeeferController
}

// The User model map to user table in db
type User struct {
	Id       int
	Username string `orm:"size(50)"`
	Email    string `orm:"size(200)"`
	Password string `orm:"size(100)"`
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
		user := &User{Username: username, Password: password}

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

func init() {
	// set default database
	orm.RegisterDataBase("default", "sqlite3", "./db/beefer.db", 30)

	// register model
	orm.RegisterModel(new(User))

	// create table
	orm.RunSyncdb("default", false, true)
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
