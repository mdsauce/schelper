package problems

import (
	"bufio"
	"bytes"
	"os"
	"strings"

	"github.com/mdsauce/schelper/logger"
)

// ReadLog will read line by line and
// analyze the strings as bytes
func ReadLog(sclog string, verbose bool) {
	AllProbs = AllProblems()
	fp, err := os.Open(sclog)
	if err != nil {
		logger.Disklog.Warnf("Could not open file %s", sclog)
		return
	}
	args := "not found"
	var lifecycle scLifecycle
	lifecycle.initLifecycle()
	lineNum := 1
	problems := make(map[string]int)
	scanner := bufio.NewScanner(fp)
	reply := makiReply()
	for scanner.Scan() {
		line := scanner.Bytes()
		if launchArgs(line) {
			a := strings.Split(string(line), " ")
			if len(a) >= 3 {
				args = strings.Join(a[3:], " ")
			}
		}
		if isProblem, problem := problem(line); isProblem == true {
			// if this is a known problem add it to the problemsdata if it isn't present
			_, present := problems[problem.Name]
			if present {
				problems[problem.Name] = problems[problem.Name] + 1
				if verbose {
					singleOutput(problem, line)
				}
			} else {
				singleOutput(problem, line)
				problems[problem.Name] = problems[problem.Name] + 1
			}
		}
		if !reply(line) {
			reply(line)
		}

		for i := range lifecycle.stages {
			if lifecycle.stages[i].reached == false && bytes.Contains(line, []byte(lifecycle.stages[i].target)) {
				lifecycle.stages[i].line = lineNum
				lifecycle.stages[i].reached = true
			}
		}
		lineNum++
	}
	logger.Disklog.Info("Tunnel Launch Arguments: ", args)
	if veryVerbose(args) {
		makiReply := reply([]byte("reply?"))
		if !makiReply {
			problems["NoMakiReply"] = problems["NoMakiReply"] + 1
		}
		noMakiReplyOutput(makiReply)
	}
	problemsOutput(problems)
	lifecycle.lifecycleOutput()
}
