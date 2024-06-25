package password_hashes

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func MD5() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Password Hash: MD5", interfaces.FindingPriorityMedium)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(^|:|;)(\\$1\\$([./A-Za-z0-9]){1,8}\\$[./A-Za-z0-9]{22})($|:|;)"), 2)
	return pattern
}
