package public_keys

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func PublicKeyInXml() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Public Key", interfaces.FindingPriorityInformative)
	pattern.AddRegexPattern(regexp.MustCompile("(?is)<RSAKeyValue>.*</RSAKeyValue>"), 0)
	return pattern
}
