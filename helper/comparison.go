package helper

import (
	"bytes"
)

func problem(logline []byte) (bool, KnownProblem) {
	if checkForEmpty(logline) {
		return false, KnownProblem{}
	}
	for _, problem := range AllProbs {
		splitline := bytes.Split(logline, []byte(" "))
		// start at (include) 3rd element in slice
		corelog := bytes.Join(splitline[3:], []byte(" "))

		headerTest := make(chan bool)
		partialTest := make(chan bool)
		go headerMatch(splitline[3:], problem.Logs, headerTest)
		go partialMatch(splitline[3:], problem.Logs, partialTest)
		headerFound := <-headerTest
		generalFound := <-partialTest

		// Lastly we do a standard Contains() check
		if bytes.Contains(problem.Logs, corelog) || headerFound || generalFound {
			return true, problem
		}
	}
	return false, KnownProblem{}
}

// headerMatch looks at the "header" of a log (First 3 strings)
// This captures unique cases where there is no timestamp or other default info
func headerMatch(logline [][]byte, problem []byte, found chan bool) {
	if len(logline) >= 3 {
		header := bytes.Join(logline[:3], []byte(" "))

		if bytes.Contains(problem, header) {
			found <- true
		}
	}
	found <- false
}

// partialMatch is a less precise method where we start at the end of a logline
// and move to the middle.  This allows exclusions of unique strings like domain names
// We call Contains against to see if shrunkLine is inside problem until we've eaten half the logline
func partialMatch(logline [][]byte, problem []byte, found chan bool) {
	for i := 0; i <= len(logline)/2; i++ {
		final := len(logline) - i
		shrunkLine := bytes.Join(logline[final-1:], []byte(" "))
		// fmt.Println(string(shrunkLine))
		if bytes.Contains(problem, shrunkLine) {
			found <- true
			return
		}
	}
	found <- false
}

// Lifecycle stuff from here on out
// --------------------------------
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
	end := len(splitline) - 1
	core := bytes.Join(splitline[3:end], []byte(" "))
	if bytes.Contains([]byte("Started scproxy on port"), core) {
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
	end := len(splitline) - 1
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:end], []byte(" "))
	if bytes.Contains([]byte("Connection closed"), core) {
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
	if bytes.Contains([]byte("Goodbye."), core) {
		return true
	}
	return false
}

func checkForEmpty(logline []byte) bool {
	splitline := bytes.Split(logline, []byte(" "))
	if len(splitline) < 4 {
		return true
	}
	return false
}
