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
	RawAnswer string
}

func getRecordType(recordtype string) uint16 {
	lookup := map[string]uint16{
		"A": dns.TypeA,
		"NS": dns.TypeNS,
		"CNAME": dns.TypeCNAME,
		"TXT": dns.TypeTXT,
		"SOA": dns.TypeSOA,
	}

	return lookup[recordtype]
}

func Query(domain string, recordtype string, nameserver string) ([]Record, time.Duration) {
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
		raw := answer.String()
		switch t := answer.(type) {
		case *dns.A:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: t.A.String(),
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "A",
				RawAnswer: raw,
			}
		case *dns.CNAME:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: t.Target,
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "CNAME",
				RawAnswer: raw,
			}
		case *dns.TXT:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: strings.Join(t.Txt, " "),
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "TXT",
				RawAnswer: raw,
			}
		case *dns.SOA:
			results[i] = Record{
				Query: t.Hdr.Name,
				Result: t.Ns + " " + t.Mbox,
				TTL: ttl,
				Nameserver: nameserver,
				RecordType: "SOA",
				RawAnswer: raw,
			}
		}
		
	}

	return results, rtt
}

