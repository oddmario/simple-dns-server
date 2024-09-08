package utils

import (
	"log"
	"mario/simple-dns-server/constants"
	"os"

	"github.com/tidwall/gjson"
)

var Config gjson.Result

var IsProcessUnstoredQueriesEnabled bool = false
var Server_ProcessUnstoredQueries string = ""

func LoadConfig() bool {
	cfg_content, _ := os.ReadFile(constants.ConfigFilePath)
	cfgContentString := BytesToString(cfg_content)

	if !gjson.Valid(cfgContentString) {
		log.Fatal("[ERROR] Malformed configuration file")
	} else {
		Config = gjson.Parse(cfgContentString)
	}

	IsProcessUnstoredQueriesEnabled = Config.Get("process_unstored_dns_queries.is_enabled").Bool()
	Server_ProcessUnstoredQueries = Config.Get("process_unstored_dns_queries.dns_server").String()

	return true
}
