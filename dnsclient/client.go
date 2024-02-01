package dnsclient

import (
	"errors"

	"github.com/miekg/dns"
)

func DnsClientUsingQuestion(q dns.Question, server string) ([]dns.RR, error) {
	m := new(dns.Msg)
	m.Compress = false
	m.RecursionAvailable = true
	m.RecursionDesired = true
	m.Question = make([]dns.Question, 1)
	m.Question[0] = q

	c := new(dns.Client)
	in, _, err := c.Exchange(m, server)
	if err != nil {
		return nil, err
	}

	answers := in.Answer

	if len(answers) <= 0 {
		return nil, errors.New("no records found")
	}

	for _, answer := range answers {
		if answer.Header().Rrtype != q.Qtype {
			return nil, errors.New("unexpected DNS type returned from the server")
		}
	}

	return answers, nil
}

func DnsClient(fqdn string, dnsType uint16, server string) ([]dns.RR, error) {
	question := dns.Question{
		Name:   dns.Fqdn(fqdn),
		Qtype:  dnsType,
		Qclass: dns.ClassINET,
	}

	return DnsClientUsingQuestion(question, server)

	/*
		Example usage of the client:

		```go
		dnsQueryAnswers, err := dnsclient.DnsClient("google.com", dns.TypeA, "8.8.8.8:53")

		if err != nil {
			// handle errors here
		}

		if t, ok := dnsQueryAnswers[0].(*dns.A); ok {
			// check if the DNS answer is castable to its correct struct
			// https://pkg.go.dev/github.com/miekg/dns#A

			fmt.Println(t.A)
		}
		```
	*/
}
