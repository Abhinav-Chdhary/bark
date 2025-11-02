package results

import (
	"fmt"
	"strings"
)

// TextFormatter formats results as human-readable text
type TextFormatter struct{}

// NewTextFormatter creates a new TextFormatter
func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

// Format formats the scan result as text
func (f *TextFormatter) Format(result *ScanResult) (string, error) {
	var sb strings.Builder

	findings := result.GetFindings()
	errors := result.GetErrors()

	if len(findings) == 0 && len(errors) == 0 {
		sb.WriteString("No BARK comments found.\n")
		return sb.String(), nil
	}

	if len(findings) > 0 {
		sb.WriteString(fmt.Sprintf("Found %d BARK comment(s):\n\n", len(findings)))
		for _, finding := range findings {
			sb.WriteString(finding.String())
			sb.WriteString("\n")
		}
	}

	if len(errors) > 0 {
		sb.WriteString(fmt.Sprintf("\nEncountered %d error(s):\n", len(errors)))
		for _, err := range errors {
			sb.WriteString(fmt.Sprintf("  - %s\n", err.Error()))
		}
	}

	return sb.String(), nil
}
