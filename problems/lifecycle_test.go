package problems

import (
	"testing"
)

func ConnClosed(t *testing.T) {
	AllProbs = AllProblems()

	target := "Connection closed"
	detect, _ := problem([]byte(target))
	if detect != true {
		t.Errorf("Did not identify: %s", target)
		t.Fail()
	}

	target = "2019-07-01 13:17:37.716 [17878] PROXY 127.0.0.1:54499 (10.113.16.24) <- cmp-wlp.cmp.sbx.zone:80 connection error: socket error: Connection refused"
	detect, _ = problem([]byte(target))
	if detect == true {
		t.Errorf("False Positive: %s", target)
		t.Fail()
	}

}
