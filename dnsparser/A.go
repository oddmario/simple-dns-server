package dnsparser

import (
	"mario/simple-dns-server/sql"
	"net"

	"github.com/miekg/dns"
)

func A(m *dns.Msg, name_dot, name_nodot string) {
	var record_id int64
	var record_type string
	var record_name string
	var record_value string
	var record_ttl int64

	stmt, err := sql.Db.Prepare("SELECT id, record_type, record_name, record_value, record_ttl FROM dns_records WHERE record_name = ? AND record_type = 'A'")
	if err != nil {
		// an error has occured while preparing the SQL statement
		return
	}
	err = stmt.QueryRow(name_nodot).Scan(&record_id, &record_type, &record_name, &record_value, &record_ttl)
	if err != nil {
		// DNS record not found in the database
		return
	}
	stmt.Close()

	r := new(dns.A)
	r.Hdr = dns.RR_Header{Name: name_dot, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: uint32(record_ttl)}
	r.A = net.ParseIP(record_value)

	m.Answer = append(m.Answer, r)
}
