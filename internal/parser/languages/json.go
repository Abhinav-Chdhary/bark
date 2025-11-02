package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-json/bindings/go"
)

func JSON() Language {
	return Language{
		Name:       "JSON",
		Extensions: []string{".json", ".jsonc"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
