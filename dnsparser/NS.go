package dnsparser

import (
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/records"

	"github.com/miekg/dns"
)

func NS(m *dns.Msg, name_dot, name_nodot string) bool {
	recordsFound, records := records.GetDNSRecords(name_nodot, "NS")

	if !recordsFound {
		// DNS record(s) was/were not found

		return false
	}

	for _, record := range records {
		r := new(dns.NS)
		r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: uint32(record.TTL)}
		r.Ns = dns.Fqdn(record.Value)

		m.Answer = append(m.Answer, r)

		if record.IsDisposable {
			db.RetriedDbExec(10, "DELETE FROM dns_records WHERE id = ?", record.ID)
		}
	}

	return true
}
