package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
)

// Language represents a programming language with its tree-sitter parser
type Language struct {
	Name       string
	Extensions []string
	Parser     *sitter.Language
	Query      string
}
