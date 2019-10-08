package problems

import (
	"strings"
)

func launchArgs(line string) bool {
	if strings.Contains(line, " -u ") || strings.Contains(line, " -k ") {
		return true
	}
	return false
}
