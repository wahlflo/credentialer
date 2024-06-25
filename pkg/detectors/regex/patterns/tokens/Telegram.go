package tokens

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func TelegramBotToken() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Telegram - Bot Authentication Token", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(bot\\d+:[a-zA-z0-9]+)(\\W|$)"), 2)
	return pattern
}
