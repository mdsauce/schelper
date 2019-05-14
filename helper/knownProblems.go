package helper

// KnownProblem contains actual log entries and
// their associated Disruption and any specific next steps
type KnownProblem struct {
	Name       string
	Disruption Disruption
	Logs       []byte
	NextSteps  string
}

// AllProblems returns all known problems
func AllProblems() []KnownProblem {
	var AllProblems []KnownProblem
	var dnsResolution = KnownProblem{Name: "DNS-Resolution", Disruption: localDNS, Logs: []byte("MAIN DNS error: non-recoverable failure in name resolution (4) MAIN DNS error: EVUTIL_EAI_FAIL MAIN DNS error"), NextSteps: `1) Locate the DNS servers that were used from the SC logs
2) See what domain name was attempting to be resolved.  Should be a 'connecting' message prior to DNS failure`}
	AllProblems = append(AllProblems, dnsResolution)

	var earlyDisconnect = KnownProblem{Name: "Early-Disconnect", Disruption: kgpConnection, Logs: []byte("KGP libevent connection error MAIN main loop exited, return code: 5"), NextSteps: `Confirm this only happens one time, at the end of the log after the Stop signal, CTRL-C (SIGINT), is sent.  If this happens during any other time there is potential for a bad connection or a Maki that has problems maintain a connection to a client.
	
	This could be the customer network having problems maintaining the TCP tunnel or problems with the Keep Alive signal.  Look for any DEAD or LIVE signals in Sumo.`}
	AllProblems = append(AllProblems, earlyDisconnect)

	var noTunnelConnection = KnownProblem{Name: "No-Initial-Tunnel", Disruption: kgpConnection, Logs: []byte("sent reply 000000000000"), NextSteps: `This may not be a problem.  It can happen a handful of times before seeing the corresponding 000000000001 sent reply.  If it happens reliably and only on the customer's network then there is a problem opening a tunnel, i.e. connecting to the Maki Tunnel VM.
	
	Confirm that you can curl -vv a maki and the subdomain.  You should also use ping to see if there is packet loss.  Networking/IT team from customers side will probably need to intervene.`}
	AllProblems = append(AllProblems, noTunnelConnection)

	return AllProblems
}
