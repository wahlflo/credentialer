package password_hashes

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func Sha256() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Password Hash: SHA256", interfaces.FindingPriorityMedium)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(^|:|;)(\\$5\\$([./A-Za-z0-9]){1,16}\\$[./A-Za-z0-9]{43})($|:|;)"), 2)
	return pattern
}
