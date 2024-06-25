package hardcoded_credentials

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func AuthorizationHttpHeader() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("HTTP Header contains credentials", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("Authorization=(basic|bearer)\\s+([a-zA-Z0-9_\\-:\\\\.=]+)"), 2)
	return pattern
}
