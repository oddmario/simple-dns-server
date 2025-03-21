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

	// https://www.alexedwards.net/blog/configuring-sqldb & https://go.dev/doc/database/manage-connections
	database.SetConnMaxLifetime(0) // default value (no expiry; this keeps the expiry up to the idle connection timeout)
	database.SetConnMaxIdleTime(time.Minute * 2)
	database.SetMaxOpenConns(int(config.Config.DbMaxOpenCons))
	database.SetMaxIdleConns(int(config.Config.DbMaxIdleCons)) // it's NEVER recommended to set this to zero (which disables idle connections overall) because idle connections help prevent some common MariaDB/MySQL issues. see https://stackoverflow.com/a/35889853/8524395

	Db = database
}
