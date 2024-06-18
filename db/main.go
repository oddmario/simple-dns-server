package db

import (
	"database/sql"
	"log"
	"mario/simple-dns-server/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string
var DB_NAME string

var Db *sql.DB

func InitDbCredentials() {
	DB_HOST = utils.Config.Get("db.host").String()
	DB_USERNAME = utils.Config.Get("db.username").String()
	DB_PASSWORD = utils.Config.Get("db.password").String()
	DB_NAME = utils.Config.Get("db.name").String()
}

func InitDb() {
	InitDbCredentials()

	database, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":3306)/"+DB_NAME)

	if err != nil {
		log.Fatal(err)
	}

	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetConnMaxIdleTime(time.Minute * 3)
	database.SetMaxOpenConns(int(utils.Config.Get("db.max_open_cons").Int()))
	database.SetMaxIdleConns(int(utils.Config.Get("db.max_idle_cons").Int()))

	Db = database
}
