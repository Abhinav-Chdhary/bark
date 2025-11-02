package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-yaml/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func YAML() Language {
	return Language{
		Name:       "YAML",
		Extensions: []string{".yml", ".yaml"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
