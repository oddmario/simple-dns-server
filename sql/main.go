package sql

import (
	"mario/simple-dns-server/utils"
)

var DB_HOST string
var DB_USERNAME string
var DB_PASSWORD string
var DB_NAME string

func InitDbCredentials() {
	DB_HOST = utils.Config.Get("db.host").String()
	DB_USERNAME = utils.Config.Get("db.username").String()
	DB_PASSWORD = utils.Config.Get("db.password").String()
	DB_NAME = utils.Config.Get("db.name").String()
}
