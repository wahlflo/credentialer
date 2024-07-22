package magic_byte_detectors

import (
	"bytes"
	"github.com/wahlflo/credentialer/pkg/interfaces"
)

type MagicByteDetector struct {
	patterns []Pattern
}

func NewMagicByteDetector() *MagicByteDetector {
	patterns := make([]Pattern, 0)

	patterns = append(patterns, &pattern{
		magicBytesInHex: "03d9a29a65fb4b",
		fileType:        "KeePass KDB Database 1.x (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "03d9a29a67fb4b",
		fileType:        "KeePass KDBX Database 2.x (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "0000000100000014",
		fileType:        "Bouncy Castle BKS V1 Database (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "0000000200000014",
		fileType:        "Bouncy Castle BKS V2 Database (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "cececece0000000200000000",
		fileType:        "JCEKS Keystore (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "feedfeed0000000200000000",
		fileType:        "JKS Keystore (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "308201963081d930819506092a8648",
		fileType:        "JBouncy Castle BCFKS Database (Encrypted Password Vault)",
	})
	patterns = append(patterns, &pattern{
		magicBytesInHex: "3056020103301106092a864886f70d",
		fileType:        "PKCS#12 File (Encrypted Password Vault)",
	})

	return &MagicByteDetector{
		patterns: patterns,
	}
}

func (detector *MagicByteDetector) AddPattern(pattern Pattern) {
	detector.patterns = append(detector.patterns, pattern)
}

func (detector *MagicByteDetector) Inject(llm interfaces.LlmConnector) {
	// LLM is not useful for this detector
}

func (detector *MagicByteDetector) Check(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile) error {
	for _, pattern := range detector.patterns {
		detector.checkPattern(output, fileToCheck, pattern)
	}
	return nil
}

func (detector *MagicByteDetector) checkPattern(output interfaces.OutputFormatter, fileToCheck interfaces.LoadedFile, pattern Pattern) {
	if bytes.HasPrefix(fileToCheck.GetContent(), pattern.GetMagicBytes()) {
		output.AddFinding(createFinding(fileToCheck, pattern.GetDetectedFileType()))
	}
}

func createFinding(fileToCheck interfaces.LoadedFile, fileType string) interfaces.Finding {
	return &interfaces.FindingInstance{
		File:                    fileToCheck,
		Name:                    "Interesting File Type: " + fileType,
		ContainsValue:           false,
		IsCompleteFileImportant: true,
		Priority:                interfaces.FindingPriorityMedium,
	}
}
