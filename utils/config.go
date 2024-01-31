package utils

import (
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
)

var Config gjson.Result

func LoadConfig() bool {
	path, _ := filepath.Abs("./config.json")
	cfg_content, _ := os.ReadFile(path)
	Config = gjson.Parse(string(cfg_content))
	return true
}
