package results

import (
	"sync"
)

// ScanResult contains all findings and errors from a scan
type ScanResult struct {
	Findings []Finding
	Errors   []error
	mu       sync.Mutex
}

// NewScanResult creates a new ScanResult
func NewScanResult() *ScanResult {
	return &ScanResult{
		Findings: make([]Finding, 0),
		Errors:   make([]error, 0),
	}
}

// AddFinding adds a finding to the result in a thread-safe manner
func (r *ScanResult) AddFinding(finding Finding) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Findings = append(r.Findings, finding)
}

// AddError adds an error to the result in a thread-safe manner
func (r *ScanResult) AddError(err error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Errors = append(r.Errors, err)
}

// HasFindings returns true if any BARK comments were found
func (r *ScanResult) HasFindings() bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return len(r.Findings) > 0
}

// GetFindings returns a copy of all findings
func (r *ScanResult) GetFindings() []Finding {
	r.mu.Lock()
	defer r.mu.Unlock()
	findings := make([]Finding, len(r.Findings))
	copy(findings, r.Findings)
	return findings
}

// GetErrors returns a copy of all errors
func (r *ScanResult) GetErrors() []error {
	r.mu.Lock()
	defer r.mu.Unlock()
	errors := make([]error, len(r.Errors))
	copy(errors, r.Errors)
	return errors
}
