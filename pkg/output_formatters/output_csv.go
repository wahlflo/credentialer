package output_formatters

import (
	"fmt"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"os"
)

type outputCsv struct {
	baseOutput
}

// NewOutputCsv generates an output formatter for csv tables
// where each finding is represented in a row in the resulting table
func NewOutputCsv(file *os.File, useColor bool) interfaces.OutputFormatter {
	output := &outputCsv{
		baseOutput: *newBaseOutput(file, useColor),
	}
	output.printoutFinding = output.PrintoutFinding
	return output
}

func (output *outputCsv) Start() {
	output.Log("priority,finding,value,location\n")
}

func (output *outputCsv) PrintoutFinding(finding interfaces.Finding) {
	findingName := fmt.Sprintf("%q", finding.GetName())

	value := ""
	if finding.GetContainsValue() {
		value = fmt.Sprintf("%q", finding.GetValue())
	}

	location := fmt.Sprintf("%q", finding.GetFile().GetFilepath())

	priority := finding.GetFindingPriority().ToString()

	output.LogWithColor(finding.GetFindingPriority(), priority+","+findingName+","+value+","+location+"\n")
}

func (output *outputCsv) Finished() {

}
