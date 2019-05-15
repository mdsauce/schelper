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
		if bytes.Contains(problem.Logs, core) {
			return true, problem
		}
	}
	return false, KnownProblem{}
}

func clientStarting(logline []byte) bool {
	if checkForEmpty(logline) {
		return false
	}
	splitline := bytes.Split(logline, []byte(" "))
	// start at (include) 3rd element in slice
	core := bytes.Join(splitline[3:], []byte(" "))
	if bytes.Contains([]byte("Using no proxy for connecting to Sauce Labs REST API."), core) {
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
	if bytes.Contains([]byte("Using no proxy for connecting to tunnel VM."), core) {
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
	if bytes.Contains([]byte("Sauce Connect is up, you may start your tests."), core) {
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
