package public_keys

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func PkcsPublicKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Public Key", interfaces.FindingPriorityInformative)
	pattern.AddRegexPattern(regexp.MustCompile("(?is)-----BEGIN.{0,20}? PUBLIC KEY-----(.*?)-----END.{0,20}? PUBLIC KEY-----"), 0)
	return pattern
}
