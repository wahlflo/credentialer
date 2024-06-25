package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func FacebookAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Facebook - Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(EAACEdEose0cBA[a-zA-Z0-9]{10,1000})(\\W|$)"), 2)
	return pattern
}
