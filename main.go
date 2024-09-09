package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"mario/simple-dns-server/config"
	"mario/simple-dns-server/constants"
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/workers"

	"github.com/miekg/dns"
)

func main() {
	// https://gist.github.com/walm/0d67b4fb2d5daf3edd4fad3e13b162cb

	args := os.Args[1:]

	if len(args) >= 1 {
		constants.ConfigFilePath = args[0]
	} else {
		constants.ConfigFilePath, _ = filepath.Abs("./config.json")
	}

	if _, err := os.Stat(constants.ConfigFilePath); errors.Is(err, os.ErrNotExist) {
		log.Fatal("The specified configuration file does not exist.")
	}

	fmt.Println("[INFO] Starting Simple DNS server v" + constants.Version + "...")

	config.LoadConfig()

	if config.Config.Mode == "db" {
		db.InitDb()
		defer db.Db.Close()
	}

	workers.Init()

	dns.HandleFunc(".", handleDnsRequest)

	var listenerData string = config.Config.ListenerAddress
	var listenerType string = config.Config.ListenerType

	server := &dns.Server{Addr: listenerData, Net: listenerType}

	log.Printf("Starting at %s\n", listenerData)

	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
