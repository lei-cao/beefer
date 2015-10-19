package main

import (
	"fmt"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const currentUserSessionKey string = "currentUser"

const (
	dbConfigPath  string = "conf/db.conf"
	dbSqlite3     string = "sqlite3"
	dbMysql       string = "mysql"
	dbPostgres    string = "postgres"
	defaultDbName string = "beefer"
	defaultDbHost string = "localhost"
)

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
	dbConfig, err := config.NewConfig("ini", dbConfigPath)
	if err != nil {
		log.Panicf("Can't find db configs under '%v'.", dbConfigPath)
	}
	db := dbConfig.String("db")
	dbUser := dbConfig.String(db + "::db_user")
	dbPassword := dbConfig.String(db + "::db_password")
	dbPort := dbConfig.String(db + "::db_port")
	dbHost := dbConfig.DefaultString(db+"::db_host", defaultDbHost)
	dbName := dbConfig.DefaultString(db+"::db_name", defaultDbName)
	// set default database
	switch db {
	case dbSqlite3:
		orm.RegisterDataBase("default", "sqlite3", dbName, 30)
	case dbMysql:
		conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPassword, dbHost, dbPort, dbName)
		orm.RegisterDataBase("default", "mysql", conn, 30)
	case dbPostgres:
		conn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)
		orm.RegisterDataBase("default", "postgres", conn, 30)
	}

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
