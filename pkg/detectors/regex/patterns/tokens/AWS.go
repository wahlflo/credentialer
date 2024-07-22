package tokens

import (
	"github.com/wahlflo/credentialer/llms"
	"github.com/wahlflo/credentialer/pkg/detectors/regex/patterns"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"log/slog"
	"regexp"
)

func AwsAccessKeyId() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("AWS - Access Key ID", interfaces.FindingPriorityHigh)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\\W|^)(A[SK]IA[a-zA-Z0-9]{16})(\\W|$)"), 2)
	return pattern
}

func AwsSecretKey() patterns.Pattern {
	pattern := patterns.NewSimpleRegexPattern("AWS - Secret Key", interfaces.FindingPriorityHigh)
	// the AWS - Secret Key is hard to match as it follows no specific pattern
	// to reduce false positives the value has to be surrounded by " or '
	// or on the right side of an assignment operator : or =
	pattern.AddRegexPattern(regexp.MustCompile("(?m)(\"[0-9a-zA-Z/+]{40}\")"), 1)
	pattern.AddRegexPattern(regexp.MustCompile("(?m)('[0-9a-zA-Z/+]{40}')"), 1)
	pattern.SetQualityCheck(qualityCheckAwsSecretKey)
	return pattern
}

func qualityCheckAwsSecretKey(originalFinding interfaces.Finding, fileToCheck interfaces.LoadedFile, llm interfaces.LlmConnector) interfaces.Finding {
	if llm == nil {
		return originalFinding
	}

	scriptExcerpt := llms.GenerateScriptExtractForLlmQuestion(string(fileToCheck.GetContent()), originalFinding.GetValue(), 200, 200)

	prompt := "Is the the value '" + originalFinding.GetValue() + "' in the following script a hardcoded AWS secret key?"
	prompt += llm.GetResponseOutputModifier()
	prompt += "The script: " + scriptExcerpt

	response, err := llm.GetBooleanResponse(prompt)
	if err != nil {
		slog.Warn("received an error when trying to get a response from the LLM:" + err.Error())
		// in case the LLM is not able to make a decision then approve the finding, to prevent that a real finding
		// gets ignored
		return originalFinding
	}

	if response {
		slog.Debug("LLM approved that '" + originalFinding.GetValue() + "' is an AWS secret key")
		return originalFinding
	} else {
		slog.Debug("LLM did NOT approved '" + originalFinding.GetValue() + "' as an AWS secret key")
		return nil
	}
}
