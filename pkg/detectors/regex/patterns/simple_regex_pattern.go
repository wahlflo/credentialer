package patterns

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

type RegexExpression struct {
	regexPattern  *regexp.Regexp
	searchedGroup int
}

type SimpleRegexPattern struct {
	RegexPattern []*RegexExpression
	PatternName  string
	Priority     interfaces.FindingPriority
	QualityCheck func(originalFinding interfaces.Finding) interfaces.Finding
}

func NewSimpleRegexPattern(patternName string, priority interfaces.FindingPriority) *SimpleRegexPattern {
	return &SimpleRegexPattern{
		RegexPattern: make([]*RegexExpression, 0),
		PatternName:  patternName,
		Priority:     priority,
		QualityCheck: nil,
	}
}

func (pattern *SimpleRegexPattern) AddRegexPattern(regexPattern *regexp.Regexp, matchingGroup int) {
	pattern.RegexPattern = append(pattern.RegexPattern, &RegexExpression{
		regexPattern:  regexPattern,
		searchedGroup: matchingGroup,
	})
}

func (pattern *SimpleRegexPattern) PerformQualityCheck(originalFinding interfaces.Finding) interfaces.Finding {
	if pattern.QualityCheck == nil {
		return originalFinding
	}
	return pattern.QualityCheck(originalFinding)
}

func (pattern *SimpleRegexPattern) SetQualityCheck(qualityCheck func(originalFinding interfaces.Finding) interfaces.Finding) {
	pattern.QualityCheck = qualityCheck
}

func (pattern *SimpleRegexPattern) GetPatternName() string {
	return pattern.PatternName
}

func (pattern *SimpleRegexPattern) GetMatches(fileName string, file []byte) []string {
	matches := make([]string, 0)
	for _, p := range pattern.RegexPattern {
		patternMatches := p.regexPattern.FindAllStringSubmatch(string(file), -1)
		for _, match := range patternMatches {
			matches = append(matches, match[p.searchedGroup])
		}
	}
	return matches
}

func (pattern *SimpleRegexPattern) GetFindingPriority() interfaces.FindingPriority {
	return pattern.Priority
}
