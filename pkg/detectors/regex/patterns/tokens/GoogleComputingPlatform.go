package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func GoogleCloudApiKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Google Cloud - API Key", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(AIza[0-9A-Za-z\\-_]{35})(\\W|$)"), 2)
	return pattern
}
