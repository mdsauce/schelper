package helper

// KnownProblem contains actual log entries and
// their associated Disruption and any specific next steps
type KnownProblem struct {
	Name       []byte
	Disruption Disruption
	Logs       []byte
	NextSteps  string
}

// AllProblems returns all known problems
func AllProblems() []KnownProblem {
	var AllProblems []KnownProblem
	var dnsResolution = KnownProblem{Name: []byte("DNS-Resolution"), Disruption: LocalDNS, Logs: []byte("MAIN DNS error: non-recoverable failure in name resolution (4) MAIN DNS error: EVUTIL_EAI_FAIL MAIN DNS error"), NextSteps: `1) Locate the DNS servers that were used from the SC logs
2) See what domain name was attempting to be resolved.  Should be a 'connecting' message prior to DNS failure`}
	AllProblems = append(AllProblems, dnsResolution)
	return AllProblems
}
