package credential_assignments

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"path/filepath"
	"regexp"
	"strings"
)

type CredentialAssignmentDetector struct {
}

func NewCredentialAssignmentDetector() *CredentialAssignmentDetector {
	return &CredentialAssignmentDetector{}
}

func (detector *CredentialAssignmentDetector) Check(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile) error {
	fileExtension := filepath.Ext(fileToCheck.GetFilename())
	for _, finding := range getFindings(fileExtension, fileToCheck.GetContent()) {
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

func getFindings(fileExtension string, content []byte) []*finding {
	findings := make([]*finding, 0)

	if fileExtension == ".sample" {
		return findings
	}

	for _, p := range getRegexPatterns() {
		patternMatches := p.regexExpression.FindAllStringSubmatch(string(content), -1)
		for _, match := range patternMatches {
			f := &finding{
				fullMatch:    match[0],
				variableName: match[p.variableGroup],
				value:        match[p.valueGroup],
				usedRegex:    p.regexExpression.String(),
			}

			if fineTuningApprovesFinding(fileExtension, f.variableName, f.value) {
				findings = append(findings, f)
			}
		}
	}

	return findings
}
