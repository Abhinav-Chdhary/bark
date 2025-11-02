package results

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestTextFormatter(t *testing.T) {
	formatter := NewTextFormatter()

	t.Run("Empty result", func(t *testing.T) {
		result := NewScanResult()
		output, err := formatter.Format(result)
		if err != nil {
			t.Fatalf("Format failed: %v", err)
		}
		if !strings.Contains(output, "No BARK comments found") {
			t.Error("Expected 'No BARK comments found' message")
		}
	})

	t.Run("With findings", func(t *testing.T) {
		result := NewScanResult()
		result.AddFinding(Finding{
			FilePath: "test.go",
			Line:     10,
			Column:   5,
			Comment:  "// BARK fix this",
		})
		result.AddFinding(Finding{
			FilePath: "main.go",
			Line:     20,
			Column:   1,
			Comment:  "// BARK: Remove debug",
		})

		output, err := formatter.Format(result)
		if err != nil {
			t.Fatalf("Format failed: %v", err)
		}

		if !strings.Contains(output, "Found 2 BARK comment(s)") {
			t.Error("Expected count of findings in output")
		}
		if !strings.Contains(output, "test.go:10:5") {
			t.Error("Expected first finding in output")
		}
		if !strings.Contains(output, "main.go:20:1") {
			t.Error("Expected second finding in output")
		}
	})
}

func TestJSONFormatter(t *testing.T) {
	formatter := NewJSONFormatter()

	t.Run("Empty result", func(t *testing.T) {
		result := NewScanResult()
		output, err := formatter.Format(result)
		if err != nil {
			t.Fatalf("Format failed: %v", err)
		}

		var jsonOutput JSONOutput
		err = json.Unmarshal([]byte(output), &jsonOutput)
		if err != nil {
			t.Fatalf("Failed to parse JSON: %v", err)
		}

		if jsonOutput.Count != 0 {
			t.Errorf("Expected count 0, got %d", jsonOutput.Count)
		}
	})

	t.Run("With findings", func(t *testing.T) {
		result := NewScanResult()
		result.AddFinding(Finding{
			FilePath: "test.go",
			Line:     10,
			Column:   5,
			Comment:  "// BARK: Fix this",
		})

		output, err := formatter.Format(result)
		if err != nil {
			t.Fatalf("Format failed: %v", err)
		}

		var jsonOutput JSONOutput
		err = json.Unmarshal([]byte(output), &jsonOutput)
		if err != nil {
			t.Fatalf("Failed to parse JSON: %v", err)
		}

		if jsonOutput.Count != 1 {
			t.Errorf("Expected count 1, got %d", jsonOutput.Count)
		}
		if len(jsonOutput.Findings) != 1 {
			t.Errorf("Expected 1 finding, got %d", len(jsonOutput.Findings))
		}
		if jsonOutput.Findings[0].FilePath != "test.go" {
			t.Errorf("Expected FilePath 'test.go', got '%s'", jsonOutput.Findings[0].FilePath)
		}
	})
}

func TestScanResultThreadSafety(t *testing.T) {
	result := NewScanResult()

	// Add findings concurrently
	done := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(n int) {
			result.AddFinding(Finding{
				FilePath: "test.go",
				Line:     uint(n),
				Column:   1,
				Comment:  "// BARK: Test",
			})
			done <- true
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}

	findings := result.GetFindings()
	if len(findings) != 10 {
		t.Errorf("Expected 10 findings, got %d", len(findings))
	}
}
