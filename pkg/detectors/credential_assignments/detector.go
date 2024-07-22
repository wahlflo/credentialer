package credential_assignments

import (
	"github.com/wahlflo/credentialer/llms"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"log/slog"
	"path/filepath"
	"regexp"
	"strings"
	"unicode/utf8"
)

type CredentialAssignmentDetector struct {
	llm interfaces.LlmConnector
}

func NewCredentialAssignmentDetector() *CredentialAssignmentDetector {
	return &CredentialAssignmentDetector{
		llm: nil,
	}
}

func (detector *CredentialAssignmentDetector) Inject(llm interfaces.LlmConnector) {
	detector.llm = llm
}

func (detector *CredentialAssignmentDetector) Check(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile) error {
	fileExtension := filepath.Ext(fileToCheck.GetFilename())
	for _, finding := range getFindings(fileExtension, fileToCheck.GetContent(), detector.llm) {
		output.AddFinding(interfaces.FindingInstance{
			File:                    fileToCheck,
			Name:                    "Credential Assignments",
			Value:                   finding.fullMatch,
			ContainsValue:           true,
			IsCompleteFileImportant: false,
			Priority:                interfaces.FindingPriorityMedium,
		})
	}

	return nil
}

type regexPattern struct {
	regexExpression *regexp.Regexp
	variableGroup   int
	valueGroup      int
}

func getRegexPatterns() []*regexPattern {
	variableNameRegex := "(" + strings.Join(interestingVariableNames, "|") + ")"

	return []*regexPattern{
		{ // 								test: asd
			regexExpression: regexp.MustCompile("(?im)([a-zA-Z0-9_]{0,25}" + variableNameRegex + "[a-zA-Z0-9_]{0,25})[ |\\t]*:[ |\\t]*([^\\s]{5,90})(\\W|$)"),
			variableGroup:   1,
			valueGroup:      3,
		},
		{ // Python / C# / Java / yml  etc. 		test = asd     AND     test = "asd"
			regexExpression: regexp.MustCompile("(?im)([a-zA-Z0-9_]{0,25}" + variableNameRegex + "[a-zA-Z0-9_]{0,25})[ |\\t]*=[ |\\t]*([^\\s]{5,90})(\\W|$)"),
			variableGroup:   1,
			valueGroup:      3,
		},
		{ // assignment in Golang: 			test := "asd"
			regexExpression: regexp.MustCompile("(?im)([a-zA-Z0-9_]{0,25}" + variableNameRegex + "[a-zA-Z0-9_]{0,25})[ |\\t]*:=[ |\\t]*(\"[^\\s]{5,90}\")(\\W|$)"),
			variableGroup:   1,
			valueGroup:      3,
		},
		{ // assignment in Rust: 			let mut s = String::from("Hallo");
			regexExpression: regexp.MustCompile("(?im)([a-zA-Z0-9_]{0,25}" + variableNameRegex + "[a-zA-Z0-9_]{0,25})[ |\\t]*=[ |\\t]*String::from\\((\"[^\\s]{5,90}\")\\)(\\W|$)"),
			variableGroup:   1,
			valueGroup:      3,
		},
		{ // assignment in R: 		 		s <- "Hallo"
			regexExpression: regexp.MustCompile("(?im)([a-zA-Z0-9_]{0,25}" + variableNameRegex + "[a-zA-Z0-9_]{0,25})[ |\\t]*<-[ |\\t]*(\"[^\\s]{5,90}\")(\\W|$)"),
			variableGroup:   1,
			valueGroup:      3,
		},
	}
}

type finding struct {
	fullMatch    string
	variableName string
	value        string
	usedRegex    string
}

func getFindings(fileExtension string, content []byte, llm interfaces.LlmConnector) []*finding {
	findings := make([]*finding, 0)

	if fileExtension == ".sample" {
		return findings
	}

	for _, p := range getRegexPatterns() {
		patternMatches := p.regexExpression.FindAllStringSubmatch(string(content), -1)
		for _, match := range patternMatches {

			// check if credential assignments contains only valid UTF8 characters
			if !utf8.ValidString(match[0]) {
				continue
			}

			f := &finding{
				fullMatch:    match[0],
				variableName: match[p.variableGroup],
				value:        match[p.valueGroup],
				usedRegex:    p.regexExpression.String(),
			}

			if !fineTuningApprovesFinding(fileExtension, f.variableName, f.value) {
				continue
			}

			if llm != nil {
				if approvedByLLM(llm, f, string(content)) {
					slog.Debug("finding: \"" + f.fullMatch + "\" was approved by LLM")
				} else {
					slog.Debug("finding: \"" + f.fullMatch + "\" was NOT approved by LLM")
					continue
				}
			}

			findings = append(findings, f)
		}
	}

	return findings
}

func approvedByLLM(llm interfaces.LlmConnector, finding *finding, fileContent string) bool {
	scriptExcerpt := llms.GenerateScriptExtractForLlmQuestion(fileContent, finding.value, 300, 300)

	prompt := "Is the value '" + finding.value + "' in the following script a hardcoded cleartext password and not a placeholder or something else?"
	prompt += llm.GetResponseOutputModifier()
	prompt += "The script: " + scriptExcerpt

	response, err := llm.GetBooleanResponse(prompt)
	if err != nil {
		slog.Warn("received an error when trying to get a response from the LLM:" + err.Error())
		// in case the LLM is not able to make a decision then approve the finding, to prevent that a real finding
		// gets ignored
		return true
	}
	return response
}
