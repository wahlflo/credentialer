package output_formatters

import (
	"encoding/json"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"os"
)

type outputJson struct {
	baseOutput
}

// NewOutputJson generates an output formatter for JSON lines
// where each finding is transformed in a JSON object
func NewOutputJson(file *os.File, useColor bool) interfaces.OutputFormatter {
	output := &outputJson{
		baseOutput: *newBaseOutput(file, useColor),
	}
	output.printoutFinding = output.PrintoutFinding
	return output
}

func (output *outputJson) Start() {

}

func (output *outputJson) PrintoutFinding(finding interfaces.Finding) {
	jsonMap := map[string]string{
		"name":      finding.GetName(),
		"file_path": finding.GetFile().GetFilepath(),
		"priority":  finding.GetFindingPriority().ToString(),
	}

	if finding.GetContainsValue() {
		jsonMap["value"] = finding.GetValue()
	}

	jsonStr, err := json.Marshal(jsonMap)
	if err != nil {
		panic(err)
	}
	output.Log(string(jsonStr) + "\n")
}

func (output *outputJson) Finished() {

}
