package command_line_parameters

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func HardcodedPasswordParameterInCurl() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Hardcoded MySql password parameter", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(curl\\s+-u\\s+(\"|'|)[a-zA-z0-9]+:\\S+(\"|'|)\\s+\\S+)(\\W|$)"), 2)
	return pattern
}
