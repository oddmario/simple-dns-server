package dnsparser

import (
	"mario/simple-dns-server/db"
	"mario/simple-dns-server/records"

	"github.com/miekg/dns"
)

func SRV(m *dns.Msg, name_dot, name_nodot string) bool {
	recordsFound, records := records.GetDNSRecord(name_nodot, "SRV")

	if !recordsFound {
		// DNS record(s) was/were not found

		return false
	}

	for _, record := range records {
		r := new(dns.SRV)
		r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeSRV, Class: dns.ClassINET, Ttl: uint32(record.TTL)}
		r.Priority = uint16(record.SRVPriority)
		r.Weight = uint16(record.SRVWeight)
		r.Port = uint16(record.SRVPort)
		r.Target = dns.Fqdn(record.SRVTarget)

		m.Answer = append(m.Answer, r)

		if record.IsDisposable {
			db.RetriedDbExec(10, "DELETE FROM dns_records WHERE id = ?", record.ID)
		}
	}

	return true
}
