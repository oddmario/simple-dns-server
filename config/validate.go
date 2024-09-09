package config

import "log"

func validateConfig() {
	if Config.Mode != "db" && Config.Mode != "static_records" {
		log.Fatal("[ERROR] Invalid operating mode. `mode` has to be either `db` or `static_records`")
	}
}
