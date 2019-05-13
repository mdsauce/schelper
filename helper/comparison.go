package helper

import (
	"bytes"
)

func problem(logline []byte) (bool, KnownProblem) {
	for _, problem := range AllProblems() {
		body := bytes.Split(logline, []byte(" "))
		core := bytes.Join(body[3:], []byte(" "))
		if bytes.Contains(problem.Logs, core) {
			return true, problem
		}
	}
	return false, KnownProblem{}
}
