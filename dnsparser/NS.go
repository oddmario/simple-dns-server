package dnsparser

import (
	"mario/simple-dns-server/sql"

	"github.com/miekg/dns"
)

func NS(m *dns.Msg, name_dot, name_nodot string) bool {
	res, err := sql.Db.Query("SELECT id, record_type, record_name, record_value, record_ttl FROM dns_records WHERE record_name = ? AND record_type = 'NS'", name_nodot)
	if err != nil {
		// an error has occured while preparing the SQL statement
		return false
	}
	defer res.Close()

	var recordsFound bool = false

	for res.Next() {
		recordsFound = true

		var record_id int64
		var record_type string
		var record_name string
		var record_value string
		var record_ttl int64

		err = res.Scan(&record_id, &record_type, &record_name, &record_value, &record_ttl)
		if err != nil {
			// an error has occured
			return false
		}

		r := new(dns.NS)
		r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeNS, Class: dns.ClassINET, Ttl: uint32(record_ttl)}
		r.Ns = dns.Fqdn(record_value)

		m.Answer = append(m.Answer, r)
	}

	if !recordsFound {
		// DNS record not found in the database
		return false
	}

	return true
}
