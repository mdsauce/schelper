package helper

import (
	"bufio"
	"os"

	"github.com/mdsauce/schelper/logger"
)

// ReadLog will read line by line and
// analyze the strings as bytes
func ReadLog(sclog string) {
	fp, err := os.Open(sclog)
	if err != nil {
		logger.Disklog.Warnf("Could not open file %s", sclog)
		return
	}
	scanner := bufio.NewScanner(fp)
	meta := make(map[string]int)
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
	}
	metaOutput(meta)
}
