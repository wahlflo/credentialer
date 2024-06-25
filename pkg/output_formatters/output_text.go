package output_formatters

import (
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"os"
)

type outputText struct {
	baseOutput
}

// NewOutputText generates an output formatter for plain text output
// this is the ideal formatter when the findings are viewed on a console
func NewOutputText(file *os.File, useColor bool) interfaces.OutputFormatter {
	output := &outputText{
		baseOutput: *newBaseOutput(file, useColor),
	}
	output.printoutFinding = output.PrintoutFinding
	return output
}

func (output *outputText) PrintoutFinding(finding interfaces.Finding) {
	output.Log("-----------------------------------" + "\n")
	output.LogWithColor(finding.GetFindingPriority(), "New Finding: "+finding.GetName()+"\n")
	output.LogWithColor(finding.GetFindingPriority(), "Priority: "+finding.GetFindingPriority().ToString()+"\n")

	if finding.GetContainsValue() {
		output.LogWithColor(finding.GetFindingPriority(), "Value: "+finding.GetValue()+"\n")
	}
	output.LogWithColor(finding.GetFindingPriority(), "Location: "+finding.GetFile().GetFilepath()+"\n")
	output.Log("-----------------------------------" + "\n")
}

func (output *outputText) Start() {

}

func (output *outputText) Finished() {

}
