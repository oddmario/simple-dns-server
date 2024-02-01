package dnsparser

import (
	"mario/simple-dns-server/sql"

	"github.com/miekg/dns"
)

func SRV(m *dns.Msg, name_dot, name_nodot string) bool {
	res, err := sql.Db.Query("SELECT id, record_type, record_name, record_value, record_ttl, srv_priority, srv_weight, srv_port, srv_target FROM dns_records WHERE record_name = ? AND record_type = 'SRV'", name_nodot)
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
		var srv_priority int64
		var srv_weight int64
		var srv_port int64
		var srv_target string

		err = res.Scan(&record_id, &record_type, &record_name, &record_value, &record_ttl, &srv_priority, &srv_weight, &srv_port, &srv_target)
		if err != nil {
			// an error has occured
			return false
		}

		r := new(dns.SRV)
		r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeSRV, Class: dns.ClassINET, Ttl: uint32(record_ttl)}
		r.Priority = uint16(srv_priority)
		r.Weight = uint16(srv_weight)
		r.Port = uint16(srv_port)
		r.Target = dns.Fqdn(srv_target)

		m.Answer = append(m.Answer, r)
	}

	if !recordsFound {
		// DNS record not found in the database
		return false
	}

	return true
}
