package password_hashes

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func Argon2() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Password Hash: Argon2", interfaces.FindingPriorityMedium)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(^|:|;)(\\$argon2(id|d|i)\\$v=[0-9]+\\$m=\\d+,t=\\d+,p=\\d+\\$[A-Za-z0-9+/=]+\\$[A-Za-z0-9+/=]+)($|:|;)"), 2)
	return pattern
}
