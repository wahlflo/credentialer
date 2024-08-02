package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func SlackOAuthBotAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Slack - OAuth v2 Bot Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(xoxb-[0-9]{11}-[0-9]{11}-[0-9a-zA-Z]{24})(\\W|$)"), 2)
	return pattern
}

func SlackOAuthUserAccessToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Slack - OAuth v2 User Access Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(xoxp-[0-9]{11}-[0-9]{11}-[0-9a-zA-Z]{24})(\\W|$)"), 2)
	return pattern
}

func SlackOAuthConfigurationToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Slack - OAuth v2 Configuration Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(xoxe.xoxp-1-[0-9a-zA-Z]{166})(\\W|$)"), 2)
	return pattern
}

func SlackOAuthRefreshToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Slack - OAuth v2 Refresh Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(xoxe-1-[0-9a-zA-Z]{147})(\\W|$)"), 2)
	return pattern
}

func SlackWebhookToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Slack - Webhook Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(T[a-zA-Z0-9_]{8}/B[a-zA-Z0-9_]{8}/[a-zA-Z0-9_]{24})(\\W|$)"), 2)
	return pattern
}
