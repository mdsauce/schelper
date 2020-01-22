package problems

import (
	"bytes"
)

func problem(logline []byte) (bool, KnownProblem) {
	if checkForEmpty(logline) || cmdLine(logline) || whitelist(logline) {
		return false, KnownProblem{}
	}
	logline = bytes.TrimSpace(logline)
	for _, problem := range AllProbs {
		splitline := bytes.Split(logline, []byte(" "))
		// start at (include) 3rd element in slice
		corelog := bytes.Join(splitline[3:], []byte(" "))

		headerFound := headerMatch(splitline[3:], problem.Logs)
		partialFound := partialMatch(splitline[3:], problem.Logs)

		// Lastly we do a standard Contains() check
		if bytes.Contains(problem.Logs, corelog) || headerFound || partialFound {
			return true, problem
		}
	}
	return false, KnownProblem{}
}

// headerMatch looks at the "header" of a log (First 3 strings)
// This captures unique cases where there is no timestamp or other default info
func headerMatch(logline [][]byte, problem []byte) bool {
	if len(logline) >= 3 {
		header := bytes.Join(logline[:3], []byte(" "))

		if bytes.Contains(problem, header) {
			return true
		}
	}
	return false
}

// partialMatch is a less precise method where we start at the end of a logline
// and move to the middle.  This allows exclusions of unique strings like domain names
// We call Contains against to see if shrunkLine is inside problem until we've eaten half the logline
func partialMatch(logline [][]byte, problem []byte) bool {
	for i := 2; i <= len(logline)/2; i++ {
		final := len(logline) - i
		shrunkLine := bytes.Join(logline[final-1:], []byte(" "))
		if bytes.Contains(problem, shrunkLine) {
			// fmt.Printf("comparing: %s <---to---> %s\n", string(problem), string(shrunkLine))
			return true
		}
	}
	return false
}

// Lifecycle stuff
// ---------------
func clientStarting(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("connecting to Sauce Labs REST API")) {
		return true
	}
	return false
}

func clientStarted(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("Started scproxy on port")) {
		return true
	}
	return false
}

func connectingToMaki(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("connecting to tunnel VM")) {
		return true
	}
	return false
}

func scUp(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("Sauce Connect is up, you may start your tests.")) {
		return true
	}
	return false
}

func scTunnelClosed(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("Connection closed")) {
		return true
	}
	return false
}

func scClientClosed(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains(core, []byte("Goodbye.")) {
		return true
	}
	return false
}

// ----End of Lifecycle stuff----

func checkForEmpty(logline []byte) bool {
	splitline := bytes.Split(logline, []byte(" "))
	if len(splitline) < 4 {
		return true
	}
	return false
}

func cmdLine(logline []byte) bool {
	if bytes.Contains(logline, []byte("Command line arguments")) {
		return true
	}
	return false
}

func whitelist(logline []byte) bool {
	if bytes.Contains(logline, []byte("MAIN created client listener on port 4445")) {
		return true
	}
	return false
}

// makiReply will return a function that can track if the
// all important maki reply 000000000001 was ever sent.  The returned
// function should be allocated to a variable.
func makiReply() func([]byte) bool {
	reply := false
	return func(logline []byte) bool {
		if reply { //skip if true
			return reply
		}
		if bytes.Contains(logline, []byte("000000000001")) {
			reply = true
		}
		return reply
	}
}
