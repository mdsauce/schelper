package problems

import (
	"fmt"
	"testing"
)

func TestEarlyDisconnect(t *testing.T) {
	AllProbs = AllProblems()
	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] KGP libevent connection"))
	if detect != true {
		t.Errorf("Did not identify: 2019-05-09 15:41:49.551 [75456] KGP libevent connection")
		t.Fail()
	}

	detect, _ = problem([]byte("2019-05-09 15:41:49.551 [75456] MAIN loop exited, return code: 5"))
	if detect != true {
		t.Errorf("Did not identify: 2019-05-09 15:41:49.551 [75456] MAIN loop exited, return code: 5")
		t.Fail()
	}

	target := "MAIN main loop exited, return code: 5"
	detect, _ = problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

}

func TestDNSResolution(t *testing.T) {
	AllProbs = AllProblems()

	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] DNS error"))
	if detect != true {
		t.Fail()
	}

	detect, _ = problem([]byte("2019-05-09 15:41:49.551 [75456] EVUTIL_EAI_FAIL MAIN DNS error"))
	if detect != true {
		t.Fail()
	}

	target := "2019-09-20 11:44:13.012 [28904] PROXY header field 'Host'"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}

}
func TestNoTunnelConn(t *testing.T) {
	AllProbs = AllProblems()

	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] 000000000000"))
	if detect != true {
		t.Errorf("Did not identify 2019-05-09 15:41:49.551 [75456] 000000000000")
		t.Fail()
	}

	detect, _ = problem([]byte("2019-05-09 15:41:49.551 [75456] CMD sent reply 000000000000"))
	if detect != true {
		t.Errorf("Did not find: CMD sent reply 000000000000")
		t.Fail()
	}
}

func TestSocketErr(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-05-15 09:01:02.291 [21134] MAIN failed to connect KGP (socket error: socket error: Connection timed out "
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify %s", target)
		t.Fail()
	}

	target = "2019-05-15 09:01:02.291 [21134] MAIN failed to connect KGP (socket error: socket error: Connection timed out)"
	detect, _ = problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify %s", target)
		t.Fail()
	}
}

func TestSSLErr(t *testing.T) {
	AllProbs = AllProblems()

	target := "20190429 141149.648 [9108] MAIN SSL verify error:num=20:unable to get local issuer certificate:depth=0:/C=US/ST=CA/O=Sauce Labs Inc/OU=Operations/CN=maki9.saucelabs.com"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-08-23 14:29:44.192 [1784] CHANNEL 2147483649 connected to ext server"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False Positive on this string: %s", target)
		t.Fail()
	}
}

func TestDNSErr(t *testing.T) {
	AllProbs = AllProblems()

	target := "20190429 141149.648 [9108] MAIN DNS error: nodename nor servname provided, or not known (-908)"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = `2018-11-02 11:46:18.996 [35616] Command line arguments: sc -u sso-toyota.tcoe-phartheeb.kandasamy -k **** -tunnel-identifier mytunnel -v --pidfile C:\temp\sc.log 
	`
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False Positive on: %s", target)
		t.Fail()
	}

	target = ` 2018-11-02 11:46:18.996 [35616] Command line arguments: sc -u sso-toyota.tcoe-phartheeb.kandasamy -k **** -tunnel-identifier mytunnel -v --pidfile C:\temp\sc.log `
	detect, prob := problem([]byte(target))
	fmt.Println(prob)
	if detect == true {
		t.Errorf("False Positive on: %s", target)
		t.Fail()
	}

	target = `2019-06-24 11:06:52.502 [54819] Sauce Connect 4.5.3, build 4602 4b3da11 `
	detect, prob = problem([]byte(target))
	fmt.Println(prob)
	if detect == true {
		t.Errorf("False Positive on: %s", target)
		t.Fail()
	}
}

func TestNoKeepalive(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-04-24 13:47:54.744 [11156] KGP warning: no keepalive ack"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-04-24 13:47:54.744 [11156] no keepalive ack"
	detect, _ = problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-04-24 13:47:54.744 [11156] no keepalive ack for 29s"
	detect, _ = problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}
}

func TestRSTByPeer(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-05-21 14:36:41.206 [5552] PROXY 127.0.0.1:44226 (172.20.43.218) <- wwwsome-website.com:443 connection error: socket error: Connection reset by peer"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-06-18 14:22:23.771 [60871] PROXY 127.0.0.1:57004 (172.20.49.172) <- www.visa.fr:443 connection error: socket error: Connection reset by peer"
	detect, _ = problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-08-08 15:20:36.709 [24824] PROXY adding ares fd 1552 events 2"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}

	target = "Where: 2019-09-20 11:44:11.688 [28904] CURL cURL: TCP_NODELAY set"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}
}

func TestFailSendHalfClose(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-06-18 14:22:23.247 [60871] PROXY 127.0.0.1:56882 failed to send half-close"
	detect, _ := problem([]byte(target))
	// fmt.Println(prob)
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-09-20 11:44:11.986 [28904] CHANNEL 2147483649 -> half-close"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False Positive on this line: %s", target)
		t.Fail()
	}
}

func TestCreateListenerFail(t *testing.T) {
	AllProbs = AllProblems()

	target := "2018-11-02 11:46:29.003 [35616] MAIN failed to create listener on port 4445"
	detect, _ := problem([]byte(target))
	// fmt.Println(prob)
	if detect != true {
		t.Errorf("Did not catch this error: %s", target)
		t.Fail()
	}

	target = "2019-08-23 14:29:42.248 [1784] MAIN created client listener on port 4445"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}
}

func TestFalsePositives(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-08-23 14:29:42.660 [1784] KGP <- last seen seq no. from announcement 0"
	detect, _ := problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}

	target = "2019-09-20 11:44:06.651 [28904] Unable to serve metrics on localhost:8888, error was: listen tcp 127.0.0.1:8888: bind: address already in use"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}

	target = "2019-08-08 15:20:36.706 [24824] PROXY parent proxy: zscaler.emirates.com, trying to resolve it"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}
}

func TestAPIRateLimit(t *testing.T) {
	AllProbs = AllProblems()

	target := "2019-10-03 09:40:54.049 [22985] error querying from https://saucelabs.com/rest/v1/max.dobeck/tunnels?full=1, error was: {\"message\": \"API rate limit exceeded for public:66.85.49.105. See rate-limiting section in our API documentation.\"}. HTTP status: 429 Unknown Status"
	detect, _ := problem([]byte(target))
	if detect == true {
		t.Errorf("False positive on this line: %s", target)
		t.Fail()
	}
}
