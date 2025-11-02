package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-zig/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func Zig() Language {
	return Language{
		Name:       "Zig",
		Extensions: []string{".zig"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
