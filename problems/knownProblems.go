package problems

// KnownProblem contains actual log entries and
// any specific next steps
type KnownProblem struct {
	Name      string
	Logs      []byte
	NextSteps string
}

// AllProbs is all known problems
var AllProbs []KnownProblem

// AllProblems returns all known problems
func AllProblems() []KnownProblem {
	var AllProblems []KnownProblem
	var dnsResolution = KnownProblem{
		Name: "DNS-Resolution",
		Logs: []byte("MAIN DNS error: non-recoverable failure in name resolution (4) MAIN DNS error: EVUTIL_EAI_FAIL MAIN DNS error"),
		NextSteps: `1) Locate the DNS servers that were used from the SC logs
2) See what domain name was attempting to be resolved.  Should be a 'connecting' message prior to DNS failure`}
	AllProblems = append(AllProblems, dnsResolution)

	var earlyDisconnect = KnownProblem{
		Name: "Early-Disconnect",
		Logs: []byte("KGP libevent connection error MAIN loop exited, return code: 5 MAIN main loop exited, return code: 5 "),
		NextSteps: `
Confirm this only happens one time, at the end of the log after the Stop signal, CTRL-C (SIGINT), is sent.  If this happens during any other time there is potential for a bad connection or a Maki that has problems maintain a connection to a client.

This could be the customer network having problems maintaining the TCP tunnel or problems with the Keep Alive signal.  Look for any DEAD or LIVE signals in Sumo.`}
	AllProblems = append(AllProblems, earlyDisconnect)

	var noTunnelConnection = KnownProblem{Name: "No-Initial-Tunnel", Logs: []byte("CMD sent reply 000000000000"), NextSteps: `
This may not be a problem.  It can happen a handful of times 
before seeing the corresponding 000000000001 sent reply.  If it happens reliably and only on the customer's network then there is a problem opening a tunnel, i.e. connecting to the Maki Tunnel VM.
	
Confirm that you can curl -vv a maki and the subdomain.  You should also use ping to see if there is packet loss.  Networking/IT team from customers side will probably need to intervene.  In the past this has been caused by not whitelisting *.miso.saucelabs.com or *.saucelabs.com.  Whitelisting *.saucelabs.com is the best option if the customer is willing.`}
	AllProblems = append(AllProblems, noTunnelConnection)

	var custSSLTLS = KnownProblem{Name: "SSL/TLS-Customer-Cert", Logs: []byte("MAIN SSL verify error MAIN SSL verify error:num=19:self signed certificate in certificate chain:depth=3:/CN= KGP SSL error: certificate verify failed in SSL routines ssl3_get_server_certificate libevent connection error"), NextSteps: `We are rejecting this because we can't verify the Self Signed cert being used is real with any of the 3rd party Certificate Authorities.

	Verify the client is not using some weird Custom Defined Self-Signed Certificate.  I.g. check the machine they're on with the customer, each OS will have a custom way of getting and verifying Certs are valid.  Try using the --capath <capath dir> flag.`}
	AllProblems = append(AllProblems, custSSLTLS)

	var sockMakiErr = KnownProblem{Name: "Socket-Maki-Conn", Logs: []byte("MAIN failed to connect KGP (socket error: socket error: No connection could be made because the target machine actively refused it.) MAIN failed to connect KGP socket error MAIN failed to connect KGP (socket error: socket error: Connection timed out) MAIN failed to connect KGP (socket error: socket error: Connection timed out "), NextSteps: `The above error message indicates that the Sauce Connect client couldn't connect to the server (Maki). The connection could be blocked by the SC client host machine's network and/or proxy firewall. You can perform the following commands from a Command Prompt/Terminal on the host machine to test the connection. Please note that the host machine is usually owned by the customers and we have to work with them to perform the commands on their machines.
	Sauce Labs utilizes several IP blocks for our services. If the IP address "66.85.49.50" used in the example is not in service, please choose another one to run the tests. For instance, you can locate the IP address used by the tunnel in the SC client log.
	
		ping 66.85.49.50
		Expected result:
		64 bytes from 66.85.49.50: icmp_seq=0 ttl=52 time=25.881 ms
		Undesired result:
		Request timeout for icmp_seq 0
	
	This is a simple network connectivity test. However, ping does not use port 443. This does not validate HTTPS connections with SSL certificates. Also, some networks may block the internet protocol (ICMP) used by ping.
	
		curl 66.85.49.50:443
		Expected result:
		curl: (52) Empty reply from server
		Undesired result:
		curl: (7) Couldn't connect to server
	
	The command may return "Empty reply from server", which is good. This means the host machine is not blocked by a proxy or firewall and the request actually left the private network.
	
		telnet maki665.miso.saucelabs.com 443
		You may perform a telnet connection test as well if you want to cover all bases. Again you just want to make sure you don't get a 4xx Blocked by proxy.
	
	The steps above are just to ensure a request can leave the private network and reach Sauce Labs servers without being blocked or filtered. If the tests show the connection was indeed blocked, the customer will need to contact their network/firewall team to whitelist our IP blocks.`}
	AllProblems = append(AllProblems, sockMakiErr)

	var noKeepalive = KnownProblem{
		Name: "No-Keepalive",
		Logs: []byte("KGP warning: no keepalive ack KGP warning: no keepalive ack for 8s"),
		NextSteps: `A keepalive is necessary to keep the tunnel open.  You should NOT be seeing this message.  Once or twice in the logs is OK but not great.  Repeated Keepalive misses can end up killing the tunnel and is usually indicative of major networking problems.  Run these sumo queries to learn more.  

Any results are bad basically.  You want the sumo search to return 0.
_sourceName="/var/local/mount/maki1234/rw/var/log/upstart/gravina.log" 
_sourceName="/var/local/mount/makiNumberHere/rw/var/log/upstart/gravina.log" LIVE
_sourceName="/var/local/mount/makiNumberHere/rw/var/log/upstart/gravina.log" DEAD`}
	AllProblems = append(AllProblems, noKeepalive)

	var noNameProvidedDNS = KnownProblem{
		Name:      "Nodename-nor-servename-DNS-Error",
		Logs:      []byte("MAIN DNS error: nodename nor servname provided, or not known (-908)"),
		NextSteps: "The hostname on one of the lines above this couldn't be resolved. Usually its a maki subdomain and the local DNS servers aren't allowed to resolve the maki saucelabs.com subdomain.  Try using a tool like Dig from the Sauce Connect host machine to resolve the offending domain name."}
	AllProblems = append(AllProblems, noNameProvidedDNS)

	var connResetByPeer = KnownProblem{
		Name: "Conn-Reset-By-Peer",
		Logs: []byte("PROXY 127.0.0.1:44226 (172.20.43.218) <- wwwsome-website.com:443 connection error: socket error: Connection reset by peer"),
		NextSteps: `This is a problem when the peer (the site you were trying to reach) closed the connection with a RST 
packet. More info here: 
https://stackoverflow.com/questions/1434451/what-does-connection-reset-by-peer-mean
“Connection reset by peer” is the TCP/IP equivalent of slamming the phone back on the hook. It’s more polite than merely not replying, leaving one hanging. But it’s not the FIN-ACK expected of the truly polite TCP/IP converseur.`}
	AllProblems = append(AllProblems, connResetByPeer)

	var failSendHalfClose = KnownProblem{
		Name:      "Fail-Send-Half-Close",
		Logs:      []byte("PROXY 127.0.0.1:56882 failed to send half-close"),
		NextSteps: "Something upstream from Sauce Connect's host machine is refusing to accept the connections from Sauce Connect.  This is specific to the KGP protocol and the actual content of the tunnel. \nThe line following this should contain a domain that was attempted to be reached.  This may have resulted in a 503 gateway Error or some other non-200 HTTP response during a test."}
	AllProblems = append(AllProblems, failSendHalfClose)

	var createListenerFailed = KnownProblem{
		Name:      "Create-Listener-Failed",
		Logs:      []byte("failed to create listener on port 4445"),
		NextSteps: "Something is using the port specified in the error message.  Read about the --se-port flag in Sauce Connect.",
	}
	AllProblems = append(AllProblems, createListenerFailed)

	var apiRateLimit = KnownProblem{
		Name:      "API-rate-limit",
		Logs:      []byte("HTTP status: 429"),
		NextSteps: "Too many HTTP requests were sent to our API.  Please wait 15-30 minutes and try again.",
	}
	AllProblems = append(AllProblems, apiRateLimit)

	return AllProblems
}
