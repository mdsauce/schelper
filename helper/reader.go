package helper

import (
	"bufio"
	"os"

	"github.com/mdsauce/schelper/logger"
)

type scLifecycle struct {
	stage   string
	reached bool
	target  string
	line    int
}

// ReadLog will read line by line and
// analyze the strings as bytes
func ReadLog(sclog string, verbose bool) {
	AllProbs = AllProblems()
	fp, err := os.Open(sclog)
	if err != nil {
		logger.Disklog.Warnf("Could not open file %s", sclog)
		return
	}
	cycle := setupLifecycle()
	lineNum := 1
	meta := make(map[string]int)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Bytes()
		if isProblem, problem := problem(line); isProblem == true {
			// if this is a known problem add it to the metadata if it isn't present
			_, present := meta[problem.Name]
			if present {
				meta[problem.Name] = meta[problem.Name] + 1
				if verbose {
					singleOutput(problem, line)
				}
			} else {
				singleOutput(problem, line)
				meta[problem.Name] = meta[problem.Name] + 1
			}
		}

		if cycle[0].reached == false && clientStarting(line) {
			cycle[0].line = lineNum
			cycle[0].reached = true
		}
		if cycle[1].reached == false && clientStarted(line) {
			cycle[1].reached = true
			cycle[1].line = lineNum
		}
		if cycle[2].reached == false && connectingToMaki(line) {
			cycle[2].reached = true
			cycle[2].line = lineNum
		}
		if cycle[3].reached == false && scUp(line) {
			cycle[3].reached = true
			cycle[3].line = lineNum
		}
		if cycle[4].reached == false && scTunnelClosed(line) {
			cycle[4].reached = true
			cycle[4].line = lineNum
		}
		if cycle[5].reached == false && scClientClosed(line) {
			cycle[5].reached = true
			cycle[5].line = lineNum
		}
		lineNum++
	}
	metaOutput(meta)
	lifecycleOutput(cycle)
}

func setupLifecycle() [6]scLifecycle {
	var lifecycle [6]scLifecycle
	lifecycle[0].stage = "Client attempted to start"
	lifecycle[0].target = "connecting to Sauce Labs REST API"

	lifecycle[1].stage = "Client started"
	lifecycle[1].target = "Started scproxy on port"

	lifecycle[2].stage = "Client attempted to connect to Maki"
	lifecycle[2].target = "connecting to tunnel VM"

	lifecycle[3].stage = "Sauce Connect Tunnel started"
	lifecycle[3].target = "Sauce Connect is up, you may start your tests."

	lifecycle[4].stage = "Client stopping attempts to reach Maki."
	lifecycle[4].target = "Connection closed"

	lifecycle[5].stage = "Client Shutdown"
	lifecycle[5].target = "Goodbye."

	for i := range lifecycle {
		lifecycle[i].reached = false
	}

	return lifecycle
}
