package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func AwsAccessKeyId() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("AWS - Access Key ID", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(A[SK]IA[a-zA-Z0-9]{16})(\\W|$)"), 2)
	return pattern
}

func AwsSecretKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("AWS - Secret Key", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)([0-9a-zA-Z/+]{40})(\\W|$)"), 2)
	return pattern
}
