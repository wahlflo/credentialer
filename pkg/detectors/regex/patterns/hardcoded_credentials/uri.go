package hardcoded_credentials

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"regexp"
)

func CredentialsInUri() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("URI containing hardcoded credentials", interfaces.FindingPriorityMedium)

	// As there a many protocols where credentials can be hardcoded in the URI
	// like: http|https|ftp|sftp|smtp|imap|pop3|ldap|telnet|ssh|mysql|postgresql|rdp|vnc|...
	// the protocol is not part of the regex expression to be as generic as possible

	pattern.AddRegexPattern(regexp.MustCompile("://(\\w|\\d)+:.*?@((\\w|\\d)+\\.)+\\w+"), 0)
	return pattern
}
