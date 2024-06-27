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
	// the AWS - Secret Key is hard to match as it follows no specific pattern
	// to reduce false positives the value has to be surrounded by " or '
	// or on the right side of an assignment operator : or =
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\"[0-9a-zA-Z/+]{40}\")"), 1)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)('[0-9a-zA-Z/+]{40}')"), 1)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)([=|:]\\s+[0-9a-zA-Z/+]{40})(\\W|$)"), 1)
	return pattern
}
