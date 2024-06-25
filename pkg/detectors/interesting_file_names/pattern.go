package interesting_file_names

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

type Pattern interface {
	GetDetectedFileType() string
	GetRegexExpression() *regexp.Regexp
	GetFindingPriority() interfaces.FindingPriority
}

type pattern struct {
	regexPattern *regexp.Regexp
	fileType     string
	priority     interfaces.FindingPriority
}

func (pattern *pattern) GetDetectedFileType() string {
	return pattern.fileType
}

func (pattern *pattern) GetRegexExpression() *regexp.Regexp {
	return pattern.regexPattern
}

func (pattern *pattern) GetFindingPriority() interfaces.FindingPriority {
	return pattern.priority
}
