package helper

import (
	"bytes"
)

func problem(logline []byte) (bool, KnownProblem) {
	if checkForEmpty(logline) {
		return false, KnownProblem{}
	}
	for _, problem := range AllProblems() {
		splitline := bytes.Split(logline, []byte(" "))
		// start at (include) 3rd element in slice
		core := bytes.Join(splitline[3:], []byte(" "))

		headerTest := make(chan bool)
		generalTest := make(chan bool)
		go headerMatch(splitline[3:], problem.Logs, headerTest)
		go nonUniqueMatch(splitline[3:], problem.Logs, generalTest)
		headerFound := <-headerTest
		generalFound := <-generalTest

		if bytes.Contains(problem.Logs, core) || headerFound || generalFound {
			return true, problem
		}
	}
	return false, KnownProblem{}
}

func headerMatch(logline [][]byte, problem []byte, found chan bool) {
	if len(logline) >= 3 {
		header := bytes.Join(logline[:3], []byte(" "))

		if bytes.Contains(problem, header) {
			found <- true
		}
	}

	found <- false
}

func nonUniqueMatch(logline [][]byte, problem []byte, found chan bool) {
	for i := 0; i < len(logline)/2; i++ {
		final := len(logline) - i
		general := bytes.Join(logline[:final], []byte(" "))
		if bytes.Contains(problem, general) {
			found <- true
			return
		}
	}

	found <- false
}

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
