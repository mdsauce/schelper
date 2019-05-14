package helper

import (
	"bytes"
)

func problem(logline []byte) (bool, KnownProblem) {
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
