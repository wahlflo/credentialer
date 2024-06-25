package detectors

import (
	"github.com/wahlflo/credentialer/pkg/detectors/credential_assignments"
	"github.com/wahlflo/credentialer/pkg/detectors/interesting_file_names"
	"github.com/wahlflo/credentialer/pkg/detectors/magic_byte_detectors"
	"github.com/wahlflo/credentialer/pkg/detectors/regex"
	"github.com/wahlflo/credentialer/pkg/interfaces"
)

func LoadBasicDetectors() []interfaces.Detector {
	detectors := make([]interfaces.Detector, 0)
	detectors = append(detectors, interesting_file_names.NewInterestingFilenameDetector())
	detectors = append(detectors, magic_byte_detectors.NewMagicByteDetector())
	detectors = append(detectors, regex.NewRegexDetector())
	detectors = append(detectors, credential_assignments.NewCredentialAssignmentDetector())
	return detectors
}
