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

	var earlyDisconnect = KnownProblem{Name: "Early-Disconnect", Disruption: kgpConnection, Logs: []byte("KGP libevent connection error MAIN loop exited, return code: 5 MAIN main loop exited, return code: 5 "), NextSteps: `
Confirm this only happens one time, at the end of the log after the Stop signal, CTRL-C (SIGINT), is sent.  If this happens during any other time there is potential for a bad connection or a Maki that has problems maintain a connection to a client.

This could be the customer network having problems maintaining the TCP tunnel or problems with the Keep Alive signal.  Look for any DEAD or LIVE signals in Sumo.`}
	AllProblems = append(AllProblems, earlyDisconnect)

	var noTunnelConnection = KnownProblem{Name: "No-Initial-Tunnel", Disruption: kgpConnection, Logs: []byte("sent reply 000000000000"), NextSteps: `
This may not be a problem.  It can happen a handful of times 
before seeing the corresponding 000000000001 sent reply.  If it happens reliably and only on the customer's network then there is a problem opening a tunnel, i.e. connecting to the Maki Tunnel VM.
	
Confirm that you can curl -vv a maki and the subdomain.  You should also use ping to see if there is packet loss.  Networking/IT team from customers side will probably need to intervene.`}
	AllProblems = append(AllProblems, noTunnelConnection)

	var custSSLTLS = KnownProblem{Name: "SSL/TLS-Customer-Cert", Disruption: sslTLS, Logs: []byte("MAIN SSL verify error MAIN SSL verify error:num=19:self signed certificate in certificate chain:depth=3:/CN= KGP SSL error: certificate verify failed in SSL routines ssl3_get_server_certificate libevent connection error"), NextSteps: `We are rejecting this because we can't verify the Self Signed cert being used is real with any of the 3rd party Certificate Authorities.

	Verify the client is not using some weird Custom Defined Self-Signed Certificate.  I.g. check the machine they're on with the customer, each OS will have a custom way of getting and verifying Certs are valid.  Try using the --capath <capath dir> flag.`}
	AllProblems = append(AllProblems, custSSLTLS)

	var sockMakiErr = KnownProblem{Name: "Socket-Maki-Conn", Logs: []byte("MAIN failed to connect KGP (socket error: socket error: No connection could be made because the target machine actively refused it.) MAIN failed to connect KGP socket error MAIN failed to connect KGP (socket error: socket error: Connection timed out) MAIN failed to connect KGP (socket error: socket error: Connection timed out "), NextSteps: `
Sauce labs lives at a select group of IP blocks covered in the whitelist.  So if we're unable to connect to a Maki or any other endpoint try pinging that endpoint from the Sauce Connect Host machine aka the host.
	
From this host machine ping 162.222.75.78.
Curl 162.222.75.78:443 as the ping protocol does not use port 443 and may not be blocked. You may get Empty reply from server which is good.  You weren't blocked by a proxy or firewall and the request actually left the private network.
Telnet as well if you want to cover all bases: telnet maki86032.miso.saucelabs.com 443.  Again you just want to make sure you don't get a 4xx Blocked by proxy blah blah blah.

All of these steps are just to ensure a request can leave the private network without being blocked or filtered.`}
	AllProblems = append(AllProblems, sockMakiErr)

	var noKeepalive = KnownProblem{Name: "No-Keepalive", Logs: []byte("KGP warning: no keepalive ack KGP warning: no keepalive ack for 8s"), NextSteps: `A keepalive is necessary to keep the tunnel open.  You should NOT be seeing this message.  Once or twice in the logs is OK but not great.  Repeated Keepalive misses can end up killing the tunnel and is usually indicative of major networking problems.  Run these sumo queries to learn more.  

Any results are bad basically.  You want the sumo search to return 0.
_sourceName="/var/local/mount/maki1234/rw/var/log/upstart/gravina.log" 
_sourceName="/var/local/mount/makiNumberHere/rw/var/log/upstart/gravina.log" LIVE
_sourceName="/var/local/mount/makiNumberHere/rw/var/log/upstart/gravina.log" DEAD`}
	AllProblems = append(AllProblems, noKeepalive)

	var noNameProvidedDNS = KnownProblem{Name: "Nodename-nor-servename-DNS-Error", Logs: []byte("MAIN DNS error: nodename nor servname provided, or not known (-908)"), NextSteps: "The hostname on one of the lines above this couldn't be resolved. Usually its a maki subdomain and the local DNS servers aren't allowed to resolve the maki saucelabs.com subdomain.  Try using a tool like Dig from the Sauce Connect host machine to resolve the offending domain name."}
	AllProblems = append(AllProblems, noNameProvidedDNS)

	return AllProblems
}
