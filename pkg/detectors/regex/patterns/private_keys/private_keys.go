package private_keys

import (
	"errors"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"golang.org/x/crypto/ssh"
	"regexp"
	"strings"
)

func PrivateKeyGeneric() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("Private Key", interfaces.FindingPriorityHigh)

	pattern.AddRegexPattern(regexp.MustCompile("(?is)-----BEGIN.{0,20}? PRIVATE KEY( BLOCK)?-----(.*?)-----END.{0,20}? PRIVATE KEY( BLOCK)?-----"), 0)

	pattern.SetQualityCheck(privateKeyQualityCheck)

	return pattern
}

func privateKeyQualityCheck(originalFinding interfaces.Finding) interfaces.Finding {

	// Lower Priority for encrypted Private Keys
	if strings.Contains(originalFinding.GetValue(), "ENCRYPTED") {
		return interfaces.FindingInstance{
			File:                    originalFinding.GetFile(),
			Name:                    "Encrypted Private Key",
			Value:                   originalFinding.GetValue(),
			ContainsValue:           true,
			IsCompleteFileImportant: false,
			Priority:                interfaces.FindingPriorityInformative,
		}
	}

	// For OpenSSH keys one can not identify if the key is encrypted or not by the identifier
	if strings.Contains(originalFinding.GetValue(), "OPENSSH") {
		if isOpenSshKeyEncrypted(originalFinding.GetValue()) {
			return interfaces.FindingInstance{
				File:                    originalFinding.GetFile(),
				Name:                    "Encrypted OPENSSH Private Key",
				Value:                   originalFinding.GetValue(),
				ContainsValue:           true,
				IsCompleteFileImportant: false,
				Priority:                interfaces.FindingPriorityInformative,
			}
		} else {
			return interfaces.FindingInstance{
				File:                    originalFinding.GetFile(),
				Name:                    "Unencrypted OPENSSH Private Key",
				Value:                   originalFinding.GetValue(),
				ContainsValue:           true,
				IsCompleteFileImportant: false,
				Priority:                interfaces.FindingPriorityHigh,
			}
		}
	}

	return originalFinding
}

func isOpenSshKeyEncrypted(keyData string) bool {
	_, err := ssh.ParseRawPrivateKey([]byte(keyData))
	if err != nil {
		var passphraseMissingError *ssh.PassphraseMissingError
		if errors.As(err, &passphraseMissingError) {
			return true
		}
		return false
	}
	return false
}
