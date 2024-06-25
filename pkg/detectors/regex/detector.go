package regex

import (
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/command_line_parameters"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/hardcoded_credentials"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/password_hashes"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/private_keys"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/public_keys"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns/tokens"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"unicode/utf8"
)

var BasicPattern = []patterns.Pattern{
	command_line_parameters.HardcodedPasswordParameterInCurl(),
	command_line_parameters.HardcodedPasswordParameterInMySqlCli(),
	command_line_parameters.HardcodedPasswordParameterInSshpass(),

	hardcoded_credentials.AuthorizationHttpHeader(),
	hardcoded_credentials.CredentialsInUri(),

	password_hashes.Argon2(),
	password_hashes.BCrypt(),
	password_hashes.MD5(),
	password_hashes.Sha256(),
	password_hashes.Sha512(),

	private_keys.PrivateKeyInXml(),
	private_keys.PrivateKeyGeneric(),
	private_keys.PuttyPrivateKeys(),

	public_keys.PublicKeyInXml(),
	public_keys.OpenSSHPublicKey(),
	public_keys.PkcsPublicKey(),

	tokens.AmazonMarketingServicesAuthToken(),
	tokens.AwsAccessKeyId(),
	tokens.AwsSecretKey(),
	tokens.FacebookAccessToken(),
	tokens.GitHubPersonalAccessToken(),
	tokens.GitHubOAuthAccessToken(),
	tokens.GitHubAppUserToServerToken(),
	tokens.GitHubServerToServerToken(),
	tokens.GitHubAppRefreshToken(),
	tokens.GitLabPersonalAccessToken(),
	tokens.GitLabJobToken(),
	tokens.GoogleCloudApiKey(),
	tokens.MailChimpAccessToken(),
	tokens.MailgunAccessToken(),
	tokens.PyPiAuthenticationToken(),
	tokens.SlackOAuthBotAccessToken(),
	tokens.SlackOAuthUserAccessToken(),
	tokens.SlackOAuthConfigurationToken(),
	tokens.SlackOAuthRefreshToken(),
	tokens.SlackWebhookToken(),
	tokens.StripeApiKey(),
	tokens.TelegramBotToken(),
	tokens.TwilioAccessToken(),
}

type Detector struct {
	patterns []patterns.Pattern
}

func NewRegexDetector() *Detector {
	return &Detector{
		patterns: BasicPattern,
	}
}

func (detector *Detector) AddPattern(pattern patterns.Pattern) {
	detector.patterns = append(detector.patterns, pattern)
}

func (detector *Detector) Check(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile) error {
	for _, pattern := range detector.patterns {
		detector.checkPattern(output, fileToCheck, pattern)
	}
	return nil
}

func (detector *Detector) checkPattern(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile, pattern patterns.Pattern) {
	for _, match := range pattern.GetMatches(fileToCheck.GetFilename(), fileToCheck.GetContent()) {
		// only allow valid UFT8 string to minimize the false positive rate
		if utf8.ValidString(match) {
			finding := detector.createFindingOnMatch(fileToCheck, pattern, match)
			finding = pattern.PerformQualityCheck(finding)
			if finding != nil {
				output.AddFinding(finding)
			}
		}
	}
}

func (detector *Detector) createFindingOnMatch(fileToCheck interfaces.LoadedFile, pattern patterns.Pattern, match string) interfaces.Finding {
	return interfaces.FindingInstance{
		File:                    fileToCheck,
		Name:                    pattern.GetPatternName(),
		Value:                   match,
		ContainsValue:           true,
		IsCompleteFileImportant: false,
		Priority:                pattern.GetFindingPriority(),
	}
}
