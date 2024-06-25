package output_formatters

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"os"
	"sync"
)

type baseOutput struct {
	mutex              sync.Mutex
	idsOfFoundFindings map[string]struct{}
	file               *os.File
	useColor           bool
	printoutFinding    func(finding interfaces.Finding)
}

func newBaseOutput(file *os.File, useColor bool) *baseOutput {
	return &baseOutput{
		mutex:              sync.Mutex{},
		idsOfFoundFindings: make(map[string]struct{}),
		file:               file,
		useColor:           useColor,
	}
}

func (output *baseOutput) AddFinding(finding interfaces.Finding) {
	output.mutex.Lock()
	defer output.mutex.Unlock()

	findingId := calculateFindingIdentifier(finding)
	if _, found := output.idsOfFoundFindings[findingId]; !found {
		output.idsOfFoundFindings[findingId] = struct{}{}
		output.printoutFinding(finding)
	}
}

func (output *baseOutput) LogWithColor(findingPriority interfaces.FindingPriority, message string) {
	sprint := output.getColorOutputFunctionFromPriority(findingPriority)
	formattedString := sprint(message)
	if _, err := output.file.WriteString(formattedString); err != nil {
		panic(err)
	}
}

func (output *baseOutput) Log(message string) {
	if _, err := output.file.WriteString(message); err != nil {
		panic(err)
	}
}

func (output *baseOutput) getColorOutputFunctionFromPriority(findingPriority interfaces.FindingPriority) func(a ...interface{}) string {
	if !output.useColor {
		return fmt.Sprint
	}

	switch findingPriority {
	case interfaces.FindingPriorityInformative:
		return color.New(color.FgGreen).Sprint
	case interfaces.FindingPriorityLow:
		return color.New(color.FgYellow).Sprint
	case interfaces.FindingPriorityMedium:
		return color.New(color.FgMagenta).Sprint
	case interfaces.FindingPriorityHigh:
		return color.New(color.FgRed).Sprint
	default:
		panic("finding priority is not covered")
	}
}
