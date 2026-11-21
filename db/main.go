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
var DB_UNIX_SOCK string

var Db *sql.DB

func InitDbCredentials() {
	DB_HOST = config.Config.DbHost
	DB_USERNAME = config.Config.DbUsername
	DB_PASSWORD = config.Config.DbPassword
	DB_NAME = config.Config.DbName
	DB_UNIX_SOCK = config.Config.DbUnixSock
}

func InitDb() {
	InitDbCredentials()

	var database *sql.DB
	var err error

	if len(DB_UNIX_SOCK) > 0 {
		database, err = sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@unix("+DB_UNIX_SOCK+")/"+DB_NAME+"?collation=utf8mb4_general_ci&autocommit=true")
	} else {
		database, err = sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+")/"+DB_NAME+"?collation=utf8mb4_general_ci&autocommit=true")
	}

	if err != nil {
		log.Fatal(err)
	}

	// https://www.alexedwards.net/blog/configuring-sqldb & https://go.dev/doc/database/manage-connections

	// This max lifetime is important! see https://github.com/go-sql-driver/MYSQL#important-settings
	database.SetConnMaxLifetime(5 * time.Minute)
	database.SetConnMaxIdleTime(90 * time.Second) // reap idle conns fast

	// it's NEVER recommended to set the max idle connections to zero (which disables idle connections overall) because idle connections help prevent some common MariaDB/MySQL issues. see https://stackoverflow.com/a/35889853/8524395
	database.SetMaxOpenConns(int(config.Config.DbMaxOpenCons))
	database.SetMaxIdleConns(int(config.Config.DbMaxIdleCons))

	Db = database
}
