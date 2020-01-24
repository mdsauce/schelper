package problems

import (
	"strings"
)

// should build a struct here

func launchArgs(line string) bool {
	if strings.Contains(line, " -u ") || strings.Contains(line, " -k ") || strings.Contains(line, " --user ") || strings.Contains(line, " --api-key ") {
		return true
	}
	return false
}

func veryVerbose(launchArgs string) bool {
	if strings.Contains(launchArgs, "--very-verbose") || strings.Contains(launchArgs, "-vv") {
		return true
	}
	if countRepeats(launchArgs, "-v") >= 2 {
		return true
	}
	return false
}

// return how many times a substring appears in a string.
// slices the source string by whitespace
func countRepeats(source string, target string) int {
	count := 0
	s := strings.Split(source, " ")
	for _, arg := range s {
		if strings.Contains(arg, target) {
			count++
		}
	}
	return count
}
