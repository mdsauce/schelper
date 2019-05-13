package helper

// KnownProblem contains actual log entries and
// their associated Disruption and any specific next steps
type KnownProblem struct {
	Name       []byte
	Disruption Disruption
	Logs       string
	NextSteps  string
}

// DNSResolution is a known problem where the local DNS
// server is found but can't resolve certain domain names for unkown reasons
var DNSResolution = KnownProblem{Name: []byte("DNS-Resolution"), Disruption: LocalDNS, Logs: "DNS error: non-recoverable failure in name resolution (4) DNS error: EVUTIL_EAI_FAIL", NextSteps: `1) Locate the DNS servers that were used from the SC logs
2) See what domain name was attempting to be resolved.  Should be a 'connecting' message prior to DNS failure`}
