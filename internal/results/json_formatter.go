package results

import (
	"encoding/json"
)

// JSONFormatter formats results as JSON
type JSONFormatter struct{}

// NewJSONFormatter creates a new JSONFormatter
func NewJSONFormatter() *JSONFormatter {
	return &JSONFormatter{}
}

// JSONOutput represents the JSON structure for output
type JSONOutput struct {
	Findings []Finding `json:"findings"`
	Errors   []string  `json:"errors,omitempty"`
	Count    int       `json:"count"`
}

// Format formats the scan result as JSON
func (f *JSONFormatter) Format(result *ScanResult) (string, error) {
	findings := result.GetFindings()
	errors := result.GetErrors()

	// Convert errors to strings
	errorStrings := make([]string, len(errors))
	for i, err := range errors {
		errorStrings[i] = err.Error()
	}

	output := JSONOutput{
		Findings: findings,
		Errors:   errorStrings,
		Count:    len(findings),
	}

	jsonBytes, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		return "", err
	}

	return string(jsonBytes), nil
}
