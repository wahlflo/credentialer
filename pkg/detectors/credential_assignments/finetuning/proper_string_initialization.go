package finetuning

import "strings"

var sourceCodeFileExtensions = map[string]struct{}{
	".go":    {},
	".php":   {},
	".c":     {},
	".cs":    {},
	".java":  {},
	".rb":    {},
	".swift": {},
	".py":    {},
	".kt":    {},
	".rs":    {},
	".r":     {},
	".scala": {},
	".hs":    {},
	".pl":    {},
	".lua":   {},
	".cpp":   {},
	".ts":    {},
	".tsx":   {},
	".js":    {},
	".jsx":   {},
	".ps1":   {},
	".xml":   {},
}

func CheckProperStringInitialization(fileExtension string, variableName string, secretValue string) bool {
	// ensure that the value in a code file starts with " or '
	if _, isCodeFile := sourceCodeFileExtensions[fileExtension]; isCodeFile {
		if !strings.HasPrefix(secretValue, "'") && !strings.HasPrefix(secretValue, "\"") {
			return false
		}
	}

	// if the password assignment was not in a code file then no ' or " wrapping is enforced
	return true
}
