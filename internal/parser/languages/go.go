package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-go/bindings/go"
)

func Go() Language {
	return Language{
		Name:       "Go",
		Extensions: []string{".go"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
