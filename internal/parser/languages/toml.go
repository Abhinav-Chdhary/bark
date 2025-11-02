package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-toml/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func TOML() Language {
	return Language{
		Name:       "TOML",
		Extensions: []string{".toml"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
