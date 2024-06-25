package private_keys

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func PrivateKeyInXml() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Private Key - RSA Key Pair in XML", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?is)<RSAKeyPair>.*</RSAKeyPair>"), 0)
	return pattern
}
