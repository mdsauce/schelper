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
	for scanner.Scan() {
		line := scanner.Bytes()
		logger.Disklog.Infof("Match for Known Problem %s", string(line))
	}
}
