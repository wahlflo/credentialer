package private_keys

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
	"strings"
)

func PuttyPrivateKeys() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Private Key - Key for Putty", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?is)PuTTY-User-Key-File-\\d: .*"), 0)
	pattern.SetQualityCheck(puttyPrivateKeyQualityCheck)
	return pattern
}

func puttyPrivateKeyQualityCheck(originalFinding interfaces.Finding) interfaces.Finding {
	if strings.Contains(originalFinding.GetValue(), "Encryption: none") {
		return originalFinding
	}

	// lower priority if private key is encrypted
	return interfaces.FindingInstance{
		File:                    originalFinding.GetFile(),
		Name:                    "Private Key - Key for Putty (encrypted)",
		Value:                   originalFinding.GetValue(),
		ContainsValue:           originalFinding.GetContainsValue(),
		IsCompleteFileImportant: originalFinding.GetIsCompleteFileImportant(),
		Priority:                interfaces.FindingPriorityInformative,
	}
}
