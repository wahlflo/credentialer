package interesting_file_names

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
)

type InterestingFilenameDetector struct {
	patterns []Pattern
}

func NewInterestingFilenameDetector() *InterestingFilenameDetector {
	return &InterestingFilenameDetector{
		patterns: loadDefaultPatterns(),
	}
}

func (detector *InterestingFilenameDetector) ClearPatterns() {
	detector.patterns = make([]Pattern, 0)
}

func (detector *InterestingFilenameDetector) AddPattern(pattern Pattern) {
	detector.patterns = append(detector.patterns, pattern)
}

func (detector *InterestingFilenameDetector) Check(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile) error {
	detectedFindings := make([]interfaces.Finding, 0)
	for _, p := range detector.patterns {
		if match := p.GetRegexExpression().Find([]byte(fileToCheck.GetFilename())); match != nil {
			detectedFindings = append(detectedFindings, createFinding(fileToCheck, "Interesting File Name - "+p.GetDetectedFileType(), p.GetFindingPriority()))
		}
	}

	if len(detectedFindings) == 0 {
		return nil
	}

	// select finding with the highest priority
	findingWithHighestPriority := interfaces.GetFindingWithHighestPriority(detectedFindings)
	output.AddFinding(findingWithHighestPriority)
	return nil
}

func createFinding(fileToCheck interfaces.LoadedFile, fileType string, priority interfaces.FindingPriority) interfaces.Finding {
	return &interfaces.FindingInstance{
		File:                    fileToCheck,
		Name:                    "Interesting File Name: " + fileType,
		ContainsValue:           false,
		IsCompleteFileImportant: true,
		Priority:                priority,
	}
}
