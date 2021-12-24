package dns

import (
	"fmt"
	"github.com/miekg/dns"
	"strings"
	"time"
)

type Record struct {
	Query      string
	Result     string
	TTL        time.Duration
	RecordType string
	RawAnswer  string
}

func getRecordType(recordtype string) uint16 {
	lookup := map[string]uint16{
		"A":     dns.TypeA,
		"NS":    dns.TypeNS,
		"PTR":    dns.TypePTR,
		"MX":    dns.TypeMX,
		"CERT":    dns.TypeCERT,
		"SRV":    dns.TypeSRV,
		"AAAA":    dns.TypeAAAA,
		"CNAME": dns.TypeCNAME,
		"TXT":   dns.TypeTXT,
		"SOA":   dns.TypeSOA,
	}

	return lookup[recordtype]
}

func createQuery(domain string, recordtype string) *dns.Msg {
	message := new(dns.Msg)
	message.Id = dns.Id()
	message.RecursionDesired = true
	message.Question = make([]dns.Question, 1)
	message.Question[0] = dns.Question{domain + ".", getRecordType(recordtype), dns.ClassINET}

	return message
}

func runQuery(message *dns.Msg, nameserver string) (*dns.Msg, time.Duration, error) {
	client := new(dns.Client)
	result, rtt, err := client.Exchange(message, nameserver+":53")

	fmt.Println(err)

	return result, rtt, err
}

func transformRecord(answer dns.RR) Record {

	ttl := time.Duration(int64(answer.Header().Ttl) * 1000 * 1000 * 1000)
	raw := answer.String()

	switch t := answer.(type) {
	case *dns.A:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.A.String(),
			TTL:        ttl,
			RecordType: "A",
			RawAnswer:  raw,
		}
	case *dns.AAAA:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.AAAA.String(),
			TTL:        ttl,
			RecordType: "AAAA",
			RawAnswer:  raw,
		}
	case *dns.PTR:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Ptr,
			TTL:        ttl,
			RecordType: "PTR",
			RawAnswer:  raw,
		}
	case *dns.MX:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Mx,
			TTL:        ttl,
			RecordType: "MX",
			RawAnswer:  raw,
		}
	case *dns.NS:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Ns,
			TTL:        ttl,
			RecordType: "NS",
			RawAnswer:  raw,
		}
	case *dns.SRV:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Target,
			TTL:        ttl,
			RecordType: "SRC",
			RawAnswer:  raw,
		}
	case *dns.CERT:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Certificate,
			TTL:        ttl,
			RecordType: "CERT",
			RawAnswer:  raw,
		}
	case *dns.CNAME:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Target,
			TTL:        ttl,
			RecordType: "CNAME",
			RawAnswer:  raw,
		}
	case *dns.TXT:
		return Record{
			Query:      t.Hdr.Name,
			Result:     strings.Join(t.Txt, " "),
			TTL:        ttl,
			RecordType: "TXT",
			RawAnswer:  raw,
		}
	case *dns.SOA:
		return Record{
			Query:      t.Hdr.Name,
			Result:     t.Ns + " " + t.Mbox,
			TTL:        ttl,
			RecordType: "SOA",
			RawAnswer:  raw,
		}
	}

	return Record{}
}

func Query(domain string, recordtype string, nameserver string) ([]Record, time.Duration) {
	query := createQuery(domain, recordtype)
	result, latency, err := runQuery(query, nameserver)

	if (err != nil) {
		fmt.Println(err, result)	
		return make([]Record, 0), latency
	}

	results := make([]Record, len(result.Answer))

	for i, answer := range result.Answer {
		results[i] = transformRecord(answer)
	}

	return results, latency
}
