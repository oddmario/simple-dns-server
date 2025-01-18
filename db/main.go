package db

import (
	"database/sql"
	"log"
	"mario/simple-dns-server/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string
var DB_NAME string

var Db *sql.DB

func InitDbCredentials() {
	DB_HOST = config.Config.DbHost
	DB_USERNAME = config.Config.DbUsername
	DB_PASSWORD = config.Config.DbPassword
	DB_NAME = config.Config.DbName
}

func InitDb() {
	InitDbCredentials()

	database, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":3306)/"+DB_NAME+"?collation=utf8mb4_general_ci&autocommit=true")

	if err != nil {
		log.Fatal(err)
	}

	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetConnMaxIdleTime(time.Minute * 3)
	database.SetMaxOpenConns(int(config.Config.DbMaxOpenCons))
	database.SetMaxIdleConns(int(config.Config.DbMaxIdleCons)) // never make this unlimited

	Db = database
}
