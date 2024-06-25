package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func MailgunAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Mailgun - Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(key-[0-9a-zA-Z]{32})(\\W|$)"), 2)
	return pattern
}
