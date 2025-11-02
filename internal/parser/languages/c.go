package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func C() Language {
	return Language{
		Name:       "C",
		Extensions: []string{".c", ".h"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
