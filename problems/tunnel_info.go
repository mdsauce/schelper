package problems

import (
	"bytes"
	"strings"
)

const (
	usernameShort = "-u"
	usernameLong  = "--user"
	keyShort      = "-k"
	keyLong       = "--api-key"
	vvShort       = "-vv"
	vvLong        = "--very-verbose"
)

var launchArgsFingerprint = [...]string{usernameShort, usernameLong, keyShort, keyLong, vvShort, vvLong}

func launchArgs(line []byte) bool {
	for _, arg := range launchArgsFingerprint {
		if bytes.Contains(line, []byte(arg)) {
			return true
		}
	}
	return false
}

func veryVerbose(launchArgs string) bool {
	if strings.Contains(launchArgs, vvShort) || strings.Contains(launchArgs, vvLong) {
		return true
	}
	if countRepeats(string(launchArgs), "-v") >= 2 {
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
