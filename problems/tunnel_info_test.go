package problems

import (
	"testing"
)

func TestDetectVeryVerbose(t *testing.T) {
	sample := "./sc -u some-user -v -k **** -s --pidfile /local/saucelabs/1.pid --logfile /local/saucelabs/logs_Shared/Automation_Tunnel_1.log --se-port 4446 -i Grainger_Automation_Tunnel --no-remove-colliding-tunnels -v"
	if !veryVerbose(sample) {
		t.Errorf("Func returned: %v. Did not find the Very Verbose setting.  There were two -v -v flags, but separated.  %s", veryVerbose(sample), sample)
		t.Fail()
	}

	sample2 := "./sc -u some-user -v -k **** -s --pidfile /local/saucelabs/1.pid --logfile /local/saucelabs/logs_Shared/Automation_Tunnel_1.log"
	if veryVerbose(sample2) {
		t.Errorf("Test returned false positive: %v.  Only 1 '-v' present. %s", veryVerbose(sample2), sample2)
	}
}
