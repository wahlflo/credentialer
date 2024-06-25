package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func GitLabPersonalAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitLab - Personal Access Tokens", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(glpat-[a-zA-Z0-9\\-]{20})(\\W|$)"), 2)
	return pattern
}

func GitLabJobToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitLab - CI/CD Job Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(gljob-[a-zA-Z0-9\\-]{20})(\\W|$)"), 2)
	return pattern
}
