package command_line_parameters

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func HardcodedPasswordParameterInMySqlCli() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Hardcoded MySql password parameter", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(mysql\\s+(-\\w\\s+[a-zA-z0-9]+\\s+)*-p\".{3,100}\")(\\W|$)"), 2)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(mysql\\s+(-\\w\\s+[a-zA-z0-9]+\\s+)*-p'.{3,100}')(\\W|$)"), 2)
	return pattern
}
