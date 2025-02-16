package main

import (
	"log"
	"mario/simple-dns-server/config"
	"mario/simple-dns-server/dnsclient"
	"mario/simple-dns-server/dnsparser"
	"net"
	"strings"

	"github.com/miekg/dns"
)

func parseQuery(m *dns.Msg, remoteAddr net.Addr) {
	for _, q := range m.Question {
		q.Name = strings.ToLower(q.Name)

		if config.Config.IsQueryLoggingEnabled {
			log.Printf("Query for %s from %s\n", q.Name, remoteAddr.String())
		}

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

		var didFindRecord bool = false

		if q.Qtype == dns.TypeA {
			didFindRecord = dnsparser.A(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if q.Qtype == dns.TypeSRV {
			didFindRecord = dnsparser.SRV(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if q.Qtype == dns.TypeCNAME {
			didFindRecord = dnsparser.CNAME(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if q.Qtype == dns.TypeNS {
			didFindRecord = dnsparser.NS(m, qNameWithLastDot, qNameWithoutLastDot)
		}

		if !didFindRecord && config.Config.IsProcessUnstoredQueriesEnabled {
			dnsClientResponse, err := dnsclient.DnsClientUsingQuestion(q, config.Config.ProcessUnstoredQueries_Server)
			if err == nil {
				m.Answer = dnsClientResponse.Answer
				m.Rcode = dnsClientResponse.Rcode
			}
		}

		if !didFindRecord && !config.Config.IsProcessUnstoredQueriesEnabled {
			m.Rcode = dns.RcodeNameError
		}

		if didFindRecord {
			m.Rcode = dns.RcodeSuccess // should be the default, but force setting it to make sure
		}
	}
}

func handleDnsRequest(w dns.ResponseWriter, r *dns.Msg) {
	defer w.Close()

	m := new(dns.Msg)
	m.SetReply(r)
	m.Compress = false

	switch r.Opcode {
	case dns.OpcodeQuery:
		parseQuery(m, w.RemoteAddr())
	}

	w.WriteMsg(m)
}
