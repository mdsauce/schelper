package helper

import (
	"bufio"
	"os"

	"github.com/mdsauce/schelper/logger"
)

type scLifecycle struct {
	stage   string
	reached bool
	line    []byte
}

// ReadLog will read line by line and
// analyze the strings as bytes
func ReadLog(sclog string) {
	fp, err := os.Open(sclog)
	if err != nil {
		logger.Disklog.Warnf("Could not open file %s", sclog)
		return
	}
	cycle := setupLifecycle()
	meta := make(map[string]int)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Bytes()
		if isProblem, problem := problem(line); isProblem == true {
			_, present := meta[problem.Name]
			if present {
				meta[problem.Name] = meta[problem.Name] + 1
			} else {
				singleOutput(problem, line)
				meta[problem.Name] = meta[problem.Name] + 1
			}
		}
		if cycle[0].reached == false && clientStarting(line) {
			cycle[0].reached = true
			cycle[0].line = line
		}
		if cycle[1].reached == false && clientStarted(line) {
			cycle[1].reached = true
			cycle[1].line = line
		}
		if cycle[2].reached == false && connectingToMaki(line) {
			cycle[2].reached = true
			cycle[2].line = line
		}
		if cycle[3].reached == false && scUp(line) {
			cycle[3].reached = true
			cycle[3].line = line
		}
		if cycle[4].reached == false && scTunnelClosed(line) {
			cycle[4].reached = true
			cycle[4].line = line
		}
		if cycle[5].reached == false && scClientClosed(line) {
			cycle[5].reached = true
			cycle[5].line = line
		}
	}
	metaOutput(meta)
	lifecycleOutput(cycle)
}

func setupLifecycle() [6]scLifecycle {
	var lifecycle [6]scLifecycle
	lifecycle[0].stage = "Client Attempted to Start"
	lifecycle[1].stage = "Client Started"
	lifecycle[2].stage = "Client Attempted to Connect to Maki"
	lifecycle[3].stage = "Sauce Connect Tunnel Started"
	lifecycle[4].stage = "Sauce Connect Tunnel Closed"
	lifecycle[5].stage = "Client Shutdown"
	for i := range lifecycle {
		lifecycle[i].reached = false
	}

	return lifecycle
}
