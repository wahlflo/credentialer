package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func PyPiAuthenticationToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("PyPi - Authentication Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(pypi-[a-zA-z0-9]+)(\\W|$)"), 2)
	return pattern
}
