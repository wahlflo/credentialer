package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func AmazonMarketingServicesAuthToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Amazon Marketing Services - Authentication Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(\\W|^)(amzn.mws.[0-9a-f]{8}-[0-9a-f]{4}-10-9a-f1{4}-[0-9a,]{4}-[0-9a-f]{12})(\\W|$)"), 2)
	return pattern
}
