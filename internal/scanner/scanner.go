package scanner

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/debkanchan/bark/internal/parser"
	"github.com/debkanchan/bark/internal/results"
)

// Scanner handles recursive file scanning
type Scanner struct {
	parser      *parser.Parser
	workerCount int
}

// NewScanner creates a new scanner with the specified number of workers
func NewScanner() *Scanner {
	// Use number of CPUs for worker count
	workerCount := runtime.NumCPU()
	if workerCount < 1 {
		workerCount = 1
	}

	return &Scanner{
		parser:      parser.NewParser(),
		workerCount: workerCount,
	}
}

// Scan recursively scans a directory for BARK comments
func (s *Scanner) Scan(rootPath string) *results.ScanResult {
	result := results.NewScanResult()

	// Collect all files to scan
	filesToScan := make([]string, 0)

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			result.AddError(fmt.Errorf("error accessing path %s: %w", path, err))
			return nil // Continue walking
		}

		// Skip directories
		if info.IsDir() {
			// Skip hidden directories and common non-source directories
			if info.Name() != "." && len(info.Name()) > 0 && info.Name()[0] == '.' {
				return filepath.SkipDir
			}
			if info.Name() == "node_modules" || info.Name() == "vendor" || info.Name() == ".git" {
				return filepath.SkipDir
			}
			return nil
		}

		filesToScan = append(filesToScan, path)

		return nil
	})

	if err != nil {
		result.AddError(fmt.Errorf("error walking directory: %w", err))
		return result
	}

	// Process files concurrently
	s.processFilesConcurrently(filesToScan, result)

	return result
}

// processFilesConcurrently processes files using a worker pool
func (s *Scanner) processFilesConcurrently(files []string, result *results.ScanResult) {
	if len(files) == 0 {
		return
	}

	// Create a channel for file paths
	fileQueue := make(chan string, len(files))

	// Fill the queue
	for _, file := range files {
		fileQueue <- file
	}
	close(fileQueue)

	// Start workers
	var wg sync.WaitGroup
	for i := 0; i < s.workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s.worker(fileQueue, result)
		}()
	}

	// Wait for all workers to finish
	wg.Wait()
}

// worker processes files from the queue
func (s *Scanner) worker(fileQueue <-chan string, result *results.ScanResult) {
	for filePath := range fileQueue {
		findings, err := s.parser.ParseFile(filePath)
		if err != nil {
			result.AddError(fmt.Errorf("error parsing %s: %w", filePath, err))
			continue
		}

		for _, finding := range findings {
			result.AddFinding(finding)
		}
	}
}
