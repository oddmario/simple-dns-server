package utils

import (
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

var Config gjson.Result

var IsProcessUnstoredQueriesEnabled bool = false
var Server_ProcessUnstoredQueries string = ""
var IsDisposableMode bool = false

func LoadConfig() bool {
	path, _ := filepath.Abs("./config.json")
	cfg_content, _ := os.ReadFile(path)
	Config = gjson.Parse(string(cfg_content))

	IsProcessUnstoredQueriesEnabled = Config.Get("process_unstored_dns_queries.is_enabled").Bool()
	Server_ProcessUnstoredQueries = Config.Get("process_unstored_dns_queries.dns_server").String()
	IsDisposableMode = Config.Get("disposable_mode").Bool()

	return true
}
