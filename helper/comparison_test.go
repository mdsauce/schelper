package helper

import (
	"testing"
)

func TestEarlyDisconnect(t *testing.T) {
	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] KGP libevent connection"))

	if detect != true {
		t.Fail()
	}

	detect, _ = problem([]byte("2019-05-09 15:41:49.551 [75456] main loop exited, return code: 5"))

	if detect != true {
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
func TestNoTunConn(t *testing.T) {
	detect, _ := problem([]byte("2019-05-09 15:41:49.551 [75456] 000000000000"))

	if detect != true {
		t.Fail()
	}
}
