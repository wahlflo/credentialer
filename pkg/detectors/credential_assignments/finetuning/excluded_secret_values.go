package finetuning

import (
	"strings"
)

func CheckForExcludedSecretValues(fileExtension string, variableName string, secretValue string) bool {
	unwrappedSecret := unwrapSecret(secretValue)

	// ignore environment placeholder like: GitHubToken: {{GITHUB_TOKEN}}
	if strings.HasPrefix(unwrappedSecret, "{{") && strings.HasSuffix(unwrappedSecret, "}}") {
		return false
	}
	// or a placeholder like:   token: ${GITHUB_TOKEN}
	if strings.HasPrefix(unwrappedSecret, "${") && strings.HasSuffix(unwrappedSecret, "}") {
		return false
	}

	// ignore URLs as secret. Credentials in URLs are detected by another detector, so we can ignore them here
	if strings.Contains(unwrappedSecret, "https://") || strings.Contains(unwrappedSecret, "http://") {
		return false
	}

	if strings.ToLower(unwrappedSecret) == "true" || strings.ToLower(unwrappedSecret) == "false" {
		return false
	}

	stopWords := []string{
		"password",
		"placeholder",
		"secret",
		"_key",
		"username",
		"getenv(",
		"passphrase",
		"keyfile",
		"dummy",
		"env.",
		"var.",
		"params.",
		"request.headers",
		"changeit",
		"changeme",
		"keystore",
	}
	if stringContainsExcludedPhrase(secretValue, stopWords) {
		return false
	}

	return true
}

func unwrapSecret(secret string) string {
	if strings.HasPrefix(secret, "\"") && strings.HasSuffix(secret, "\"") {
		return secret[1 : len(secret)-1]
	}
	if strings.HasPrefix(secret, "'") && strings.HasSuffix(secret, "'") {
		return secret[1 : len(secret)-1]
	}
	return secret
}

func stringContainsExcludedPhrase(value string, excludedPhrases []string) bool {
	foundValueInLowerCase := strings.ToLower(value)
	for _, stopWord := range excludedPhrases {
		if strings.Contains(foundValueInLowerCase, stopWord) {
			return true
		}
	}
	return false
}
