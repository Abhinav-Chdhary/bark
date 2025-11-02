package scanner

import (
	"os"
	"path/filepath"
	"testing"
)

func TestScanner(t *testing.T) {
	scanner := NewScanner()

	t.Run("Scan directory with BARK comments", func(t *testing.T) {
		// Create test directory structure
		tmpDir := t.TempDir()

		// Create test files
		testFiles := map[string]string{
			"test1.go": `package main
// BARK: Remove this
func main() {}`,
			"test2.js": `// BARK: Fix this
console.log("test");`,
			"clean.py": `# Regular comment
def hello():
    pass`,
			"nested/test3.go": `package nested
// BARK: Nested comment
func test() {}`,
		}

		for path, content := range testFiles {
			fullPath := filepath.Join(tmpDir, path)
			err := os.MkdirAll(filepath.Dir(fullPath), 0755)
			if err != nil {
				t.Fatalf("Failed to create directory: %v", err)
			}
			err = os.WriteFile(fullPath, []byte(content), 0644)
			if err != nil {
				t.Fatalf("Failed to create file: %v", err)
			}
		}

		// Scan directory
		result := scanner.Scan(tmpDir)

		// Should find 3 BARK comments (test1.go, test2.js, nested/test3.go)
		findings := result.GetFindings()
		if len(findings) != 3 {
			t.Errorf("Expected 3 findings, got %d", len(findings))
		}

		errors := result.GetErrors()
		if len(errors) > 0 {
			t.Errorf("Expected no errors, got %d: %v", len(errors), errors)
		}
	})

	t.Run("Skip hidden directories", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create hidden directory with BARK comment
		hiddenDir := filepath.Join(tmpDir, ".hidden")
		err := os.MkdirAll(hiddenDir, 0755)
		if err != nil {
			t.Fatalf("Failed to create hidden directory: %v", err)
		}

		hiddenFile := filepath.Join(hiddenDir, "test.go")
		err = os.WriteFile(hiddenFile, []byte("package main\n// BARK: Should not be found"), 0644)
		if err != nil {
			t.Fatalf("Failed to create file: %v", err)
		}

		// Scan directory
		result := scanner.Scan(tmpDir)

		// Should find 0 BARK comments (hidden directory should be skipped)
		findings := result.GetFindings()
		if len(findings) != 0 {
			t.Errorf("Expected 0 findings (hidden dir should be skipped), got %d", len(findings))
		}
	})

	t.Run("Skip vendor and node_modules", func(t *testing.T) {
		tmpDir := t.TempDir()

		// Create vendor and node_modules directories
		for _, dir := range []string{"vendor", "node_modules"} {
			dirPath := filepath.Join(tmpDir, dir)
			err := os.MkdirAll(dirPath, 0755)
			if err != nil {
				t.Fatalf("Failed to create %s directory: %v", dir, err)
			}

			filePath := filepath.Join(dirPath, "test.go")
			err = os.WriteFile(filePath, []byte("package main\n// BARK: Should not be found"), 0644)
			if err != nil {
				t.Fatalf("Failed to create file: %v", err)
			}
		}

		// Scan directory
		result := scanner.Scan(tmpDir)

		// Should find 0 BARK comments
		findings := result.GetFindings()
		if len(findings) != 0 {
			t.Errorf(
				"Expected 0 findings (vendor/node_modules should be skipped), got %d",
				len(findings),
			)
		}
	})
}

func TestWorkerCount(t *testing.T) {
	scanner := NewScanner()
	if scanner.workerCount < 1 {
		t.Error("Worker count should be at least 1")
	}
}
