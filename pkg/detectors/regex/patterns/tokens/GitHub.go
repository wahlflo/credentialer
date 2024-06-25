package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func GitHubPersonalAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitHub - Personal Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(ghp_[a-zA-z0-9_]{36})(\\W|$)"), 2)
	return pattern
}

func GitHubOAuthAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitHub - OAuth Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(gho_[a-zA-z0-9_]{36})(\\W|$)"), 2)
	return pattern
}

func GitHubAppUserToServerToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitHub - App user-to-server Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(ghu_[a-zA-z0-9_]{36})(\\W|$)"), 2)
	return pattern
}

func GitHubServerToServerToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitHub - App server-to-server Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(ghs_[a-zA-z0-9_]{36})(\\W|$)"), 2)
	return pattern
}

func GitHubAppRefreshToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("GitHub - App refresh token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(ghr_[a-zA-z0-9_]{36})(\\W|$)"), 2)
	return pattern
}
