package output_formatters

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"os"
)

func calculateFindingIdentifier(finding interfaces.Finding) string {
	var buffer bytes.Buffer
	buffer.WriteString(finding.GetName())
	buffer.WriteString(finding.GetFile().GetFilepath())
	if finding.GetContainsValue() {
		buffer.WriteString(finding.GetValue())
	}
	return buffer.String()
}

type writer struct {
	useColor bool
	file     *os.File
}

func newWriter(useColor bool, file *os.File) *writer {
	return &writer{
		useColor: useColor,
		file:     file,
	}
}

func (w *writer) writeLine(color *color.Color, message string) {
	if w.useColor {
		w.writeLineWithColor(color, message)
	} else {
		w.writeLineWithoutColor(message)
	}
}

func (w *writer) writeLineWithColor(c *color.Color, message string) {
	c.SetWriter(w.file)
	_, _ = c.Print(message)
}

func (w *writer) writeLineWithoutColor(message string) {
	_, _ = w.file.WriteString(message)
}
