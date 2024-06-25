package public_keys

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func OpenSSHPublicKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Public Key", interfaces.FindingPriorityInformative)
	pattern.AddRegexPattern(regexp.MustCompile("(?is)ssh-(ed\\d+|rsa|ecdsa-.?) AAAA(\\w|/|\\+)+ (.{0,40})?"), 0)
	return pattern
}
