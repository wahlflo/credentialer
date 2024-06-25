package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func MailChimpAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("MailChimp - Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)([0-9a-f]{32}-us[0-9]{1,2})(\\W|$)"), 2)
	return pattern
}
