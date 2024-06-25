package password_hashes

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func Sha512() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Password Hash: SHA512", interfaces.FindingPriorityMedium)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(^|:|;)(\\$6\\$([./A-Za-z0-9]){1,16}\\$[./A-Za-z0-9]{86})($|:|;)"), 2)
	return pattern
}
