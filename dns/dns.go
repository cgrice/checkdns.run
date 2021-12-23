package dns

import (
	// "fmt"
	"strings"
	"time"
	"github.com/miekg/dns"
)

type Record struct {
	Query string
	Result string
	TTL time.Duration
	Nameserver string
	RecordType string
	Latency time.Duration
}

func getRecordType(recordtype string) uint16 {
	lookup := map[string]uint16{
		"A": dns.TypeA,
		"NS": dns.TypeNS,
		"CNAME": dns.TypeCNAME,
		"TXT": dns.TypeTXT,
	}

	return lookup[recordtype]
}

func Query(domain string, recordtype string, nameserver string) []Record {
	message := new(dns.Msg)
	message.Id = dns.Id()
	message.RecursionDesired = true
	message.Question = make([]dns.Question, 1)
	message.Question[0] = dns.Question{domain + ".", getRecordType(recordtype), dns.ClassINET}

	client := new(dns.Client)
	result, rtt, _ := client.Exchange(message, nameserver + ":53")

	results := make([]Record, len(result.Answer))

	for i, answer := range result.Answer {
		ttl := time.Duration(int64(answer.Header().Ttl) * 1000 * 1000 * 1000)
		switch t := answer.(type) {
		case *dns.A:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: t.A.String(),
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "A",
				Latency: rtt,
			}
		case *dns.CNAME:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: t.Target,
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "CNAME",
				Latency: rtt,
			}
		case *dns.TXT:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: strings.Join(t.Txt, " "),
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "TXT",
				Latency: rtt,
			}
		}
	}

	return results
}

