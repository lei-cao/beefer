package app

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
)

const (
	dbConfigPath  string = "conf/db.conf"
	dbSqlite3     string = "sqlite3"
	dbMysql       string = "mysql"
	dbPostgres    string = "postgres"
	defaultDbName string = "beefer"
	defaultDbHost string = "localhost"
)

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

	// Init User model
	Init()

	// Sync db
	orm.RunSyncdb("default", false, true)
}
