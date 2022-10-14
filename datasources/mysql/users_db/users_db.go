package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "root"
	mysql_users_password = "12345678"
	mysql_users_host     = "127.0.0.1:3306"
	mysql_users_schema   = "users_db"
)

var (
	Client   *sql.DB
	username = mysql_users_username
	password = mysql_users_password
	host     = mysql_users_host
	schema   = mysql_users_schema
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", //username , passward, @tcp( host )/ connect scheme
		username,
		password,
		host,
		schema,
	)
	// don't log any config of database
	var err error
	Client, err = sql.Open("mysql", dataSourceName) //connect
	if err != nil {                                 // fail open database
		panic(err)
	}
	if err = Client.Ping(); err != nil { //database connection fail
		panic(err)
	}

	log.Println("database sucsessfully configured")
}
