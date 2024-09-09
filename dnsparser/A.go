package dnsparser

import (
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/records"
	"net"

	"github.com/miekg/dns"
)

func A(m *dns.Msg, name_dot, name_nodot string) bool {
	recordsFound, records := records.GetDNSRecords(name_nodot, "A")

	if !recordsFound {
		// DNS record(s) was/were not found

		return false
	}

	for _, record := range records {
		r := new(dns.A)
		r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(record.TTL)}
		r.A = net.ParseIP(record.Value)

		m.Answer = append(m.Answer, r)

		if record.IsDisposable {
			db.RetriedDbExec(10, "DELETE FROM dns_records WHERE id = ?", record.ID)
		}
	}

	return true
}
