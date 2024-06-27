package credential_assignments

import "github.com/wahlflo/credentialer/pkg/detectors/credential_assignments/finetuning"

func fineTuningApprovesFinding(fileExtension string, variableName string, secretValue string) bool {
	fineTuningApprover := []func(string, string, string) bool{
		finetuning.CheckForExcludedVariableNames,
		finetuning.CheckForExcludedSecretValues,
		finetuning.CheckProperStringInitialization,
		finetuning.CheckProperCharacterFrequencyDistribution,
	}

	for _, approver := range fineTuningApprover {
		if !approver(fileExtension, variableName, secretValue) {
			return false
		}
	}
	return true
}
