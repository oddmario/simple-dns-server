package main

import (
	"log"
	"net"
	"strings"

	"mario/simple-dns-server/dnsparser"
	"mario/simple-dns-server/sql"
	"mario/simple-dns-server/utils"

	"github.com/miekg/dns"
)

func parseQuery(m *dns.Msg, remoteAddr net.Addr) {
	for _, q := range m.Question {
		log.Printf("Query for %s from %s\n", q.Name, remoteAddr.String())

		lastCharOfQNAME := q.Name[len(q.Name)-1:]

		var qNameWithLastDot string = ""
		var qNameWithoutLastDot string = ""

		if lastCharOfQNAME == "." {
			qNameWithLastDot = q.Name
			qNameWithoutLastDot = strings.TrimSuffix(q.Name, ".")
		} else {
			qNameWithLastDot = q.Name + "."
			qNameWithoutLastDot = q.Name
		}

		if q.Qtype == dns.TypeA {
			dnsparser.A(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if q.Qtype == dns.TypeSRV {
			dnsparser.SRV(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if q.Qtype == dns.TypeCNAME {
			dnsparser.CNAME(m, qNameWithLastDot, qNameWithoutLastDot)
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m, w.RemoteAddr())
	}

	w.WriteMsg(m)
}

func main() {
	// https://gist.github.com/walm/0d67b4fb2d5daf3edd4fad3e13b162cb

	utils.LoadConfig()
	sql.InitDb()

	dns.HandleFunc(".", handleDnsRequest)

	var listenerData string = utils.Config.Get("listener.data").String()
	var listenerType string = utils.Config.Get("listener.type").String()

	server := &dns.Server{Addr: listenerData, Net: listenerType}
	log.Printf("Starting at %s\n", listenerData)
	err := server.ListenAndServe()
	defer server.Shutdown()
	if err != nil {
		log.Fatalf("Failed to start server: %s\n ", err.Error())
	}
}
