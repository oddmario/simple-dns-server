package config

import (
	"log"
	"mario/simple-dns-server/constants"
	"mario/simple-dns-server/models"
	"mario/simple-dns-server/utils"
	"os"

	"github.com/tidwall/gjson"
)

var config gjson.Result
var Config *models.ConfigObject = &models.ConfigObject{
	Mode:                            "",
	DbHost:                          "",
	DbUsername:                      "",
	DbPassword:                      "",
	DbName:                          "",
	DbMaxOpenCons:                   0,
	DbMaxIdleCons:                   0,
	ListenerType:                    "",
	ListenerAddress:                 "",
	IsProcessUnstoredQueriesEnabled: false,
	ProcessUnstoredQueries_Server:   "",
	StaticRecords:                   []*models.Record{},
}

func LoadConfig() {
	cfg_content, _ := os.ReadFile(constants.ConfigFilePath)
	cfgContentString := utils.BytesToString(cfg_content)

	if !gjson.Valid(cfgContentString) {
		log.Fatal("[ERROR] Malformed configuration file")
	} else {
		config = gjson.Parse(cfgContentString)
		storeEssentialConfigValues()
	}

	validateConfig()
}

func storeEssentialConfigValues() {
	Config.Mode = config.Get("mode").String()

	Config.DbHost = config.Get("db.host").String()
	Config.DbUsername = config.Get("db.username").String()
	Config.DbPassword = config.Get("db.password").String()
	Config.DbName = config.Get("db.name").String()
	Config.DbMaxOpenCons = config.Get("db.max_open_cons").Int()
	Config.DbMaxIdleCons = config.Get("db.max_idle_cons").Int()

	Config.ListenerType = config.Get("listener.type").String()
	Config.ListenerAddress = config.Get("listener.data").String()

	Config.IsProcessUnstoredQueriesEnabled = config.Get("process_unstored_dns_queries.is_enabled").Bool()
	Config.ProcessUnstoredQueries_Server = config.Get("process_unstored_dns_queries.dns_server").String()

	Config.StaticRecords = []*models.Record{}
	config.Get("static_records").ForEach(func(key, value gjson.Result) bool {
		var record_type string = value.Get("type").String()
		var record_name string = value.Get("name").String()
		var record_value string = value.Get("value").String()
		var record_ttl int64 = value.Get("ttl").Int()
		var srv_priority int64 = value.Get("srv_priority").Int()
		var srv_weight int64 = value.Get("srv_weight").Int()
		var srv_port int64 = value.Get("srv_port").Int()
		var srv_target string = value.Get("srv_target").String()

		Config.StaticRecords = append(Config.StaticRecords, &models.Record{
			ID:           -1,
			Type:         record_type,
			Name:         record_name,
			Value:        record_value,
			TTL:          record_ttl,
			SRVPriority:  srv_priority,
			SRVWeight:    srv_weight,
			SRVPort:      srv_port,
			SRVTarget:    srv_target,
			IsDisposable: false,
		})

		return true
	})
}
