package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Language represents a programming language with its tree-sitter parser
type Language struct {
	Name             string
	Extensions       []string
	FilenamePatterns []string // Optional: regex patterns for matching filenames (e.g., `^\.env(\.|$)`)
	Parser           *sitter.Language
	Query            string
}
