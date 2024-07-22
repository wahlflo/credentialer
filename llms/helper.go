package llms

import (
	"strings"
)

// GenerateScriptExtractForLlmQuestion extracts an excerpt from a bigger file content
// this is useful to limit the input for LLMs to faster the result generation
func GenerateScriptExtractForLlmQuestion(fullFile string, findingValue string, numberOfCharsBefore int, numberOfCharsAfter int) string {
	matchIndex := strings.Index(fullFile, findingValue)
	if matchIndex == -1 {
		panic("Could not find the match for " + findingValue)
	}

	beforeMatchIndex := max(0, matchIndex-numberOfCharsBefore)
	afterMatchIndex := min(len(fullFile), matchIndex+len(findingValue)+numberOfCharsAfter)

	return fullFile[beforeMatchIndex:afterMatchIndex]
}
