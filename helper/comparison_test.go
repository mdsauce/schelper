package helper

import (
	"testing"
)

func TestEarlyDisconnect(t *testing.T) {
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
	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] DNS error"))
	if detect != true {
		t.Fail()
	}

	detect, _ = problem([]byte("2019-05-09 15:41:49.551 [75456] EVUTIL_EAI_FAIL MAIN DNS error"))
	if detect != true {
		t.Fail()
	}

}
func TestNoTunnelConn(t *testing.T) {
	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] 000000000000"))
	if detect != true {
		t.Errorf("Did not identify 2019-05-09 15:41:49.551 [75456] 000000000000")
		t.Fail()
	}
}

func TestSocketErr(t *testing.T) {
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
	target := "20190429 141149.648 [9108] MAIN SSL verify error:num=20:unable to get local issuer certificate:depth=0:/C=US/ST=CA/O=Sauce Labs Inc/OU=Operations/CN=maki9.saucelabs.com"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}
}

func TestDNSErr(t *testing.T) {
	target := "20190429 141149.648 [9108] MAIN DNS error: nodename nor servname provided, or not known (-908)"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}
}

func TestNoKeepalive(t *testing.T) {
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
