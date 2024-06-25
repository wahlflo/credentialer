package patterns

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
)

type Pattern interface {
	GetPatternName() string
	GetMatches(filename string, file []byte) []string
	GetFindingPriority() interfaces.FindingPriority
	PerformQualityCheck(originalFinding interfaces.Finding) interfaces.Finding
}
