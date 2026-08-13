package hash

import (
	"regexp"
)

var (
	HASH_REGEX = regexp.MustCompile(`^[a-fA-F0-9]{64}$`)
)

func IsValid(hash string) bool {
	if len(hash) != 64 {
		return false
	}

	if !HASH_REGEX.MatchString(hash) {
		return false
	}

	return true
}
