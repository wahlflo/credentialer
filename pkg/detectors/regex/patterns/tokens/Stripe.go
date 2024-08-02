package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func StripeApiKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Stripe - API Key", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(sk_live_[0-9a-zA-Z]{24})(\\W|$)"), 2)
	return pattern
}
