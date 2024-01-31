package sql

import (
	"database/sql"
	"mario/simple-dns-server/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDb() bool {
	InitDbCredentials()

	database, err := sql.Open("mysql", DB_USERNAME+":"+DB_PASSWORD+"@tcp("+DB_HOST+":3306)/"+DB_NAME)
	if err != nil {
		panic(err)
	}
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetMaxOpenConns(int(utils.Config.Get("db.max_open_cons").Int()))
	database.SetMaxIdleConns(int(utils.Config.Get("db.max_idle_cons").Int()))
	Db = database
	return true
}
