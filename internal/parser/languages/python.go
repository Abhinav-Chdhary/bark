package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-python/bindings/go"
)

func Python() Language {
	return Language{
		Name:       "Python",
		Extensions: []string{".py", ".pyw"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
