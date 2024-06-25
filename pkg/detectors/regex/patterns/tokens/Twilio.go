package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func TwilioAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Twilio - Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(55[0-9a-fA-F]{32})(\\W|$)"), 2)
	return pattern
}
