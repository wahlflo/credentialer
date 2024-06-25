package password_hashes

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func BCrypt() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Password Hash: BCrypt", interfaces.FindingPriorityMedium)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(^|:|;)(\\$2[ayb]\\$\\d{1,2}\\$[./A-Za-z0-9]{53})($|:|;)"), 2)
	return pattern
}
