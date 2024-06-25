package command_line_parameters

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func HardcodedPasswordParameterInSshpass() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Hardcoded sshpass password parameter", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(sshpass\\s+-p\\s+\".{3,250}\")(\\W|$)"), 2)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(sshpass\\s+-p\\s+'.{3,250}')(\\W|$)"), 2)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(sshpass\\s+-p\\s[^\"']\\S{3,250})(\\W|$)"), 2)
	return pattern
}
