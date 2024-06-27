package finetuning

import (
	"strings"
)

func CheckForExcludedVariableNames(fileExtension string, variableName string, secretValue string) bool {
	// ignore the finding if the variable is only "key" without any pre- or suffix
	ignoredVariableNames := map[string]struct{}{
		"keyref":       {},
		"key":          {},
		"itemkey":      {},
		"keyname":      {},
		"passwordhint": {},
	}
	if _, found := ignoredVariableNames[strings.ToLower(variableName)]; found {
		return false
	}

	stopWords := []string{
		"keyword",
		"public",
		"tokenizer",
		"policies",
		"policy",
		"rule",
		"secret_name",
		"secretname",
		"compass",
		"bypass",
	}
	if stringContainsExcludedPhrase(variableName, stopWords) {
		return false
	}

	return true
}
