package results

import (
	"fmt"
)

// Finding represents a single BARK comment found in a file
type Finding struct {
	FilePath string `json:"file_path"`
	Line     uint   `json:"line"`
	Column   uint   `json:"column"`
	Comment  string `json:"comment"`
}

// String returns a human-readable representation of the finding
func (f Finding) String() string {
	return fmt.Sprintf("%s:%d:%d: %s", f.FilePath, f.Line, f.Column, f.Comment)
}
