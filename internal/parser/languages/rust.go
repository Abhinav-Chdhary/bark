package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-rust/bindings/go"
)

func Rust() Language {
	return Language{
		Name:       "Rust",
		Extensions: []string{".rs"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "([(line_comment) (block_comment)] @comment)",
	}
}
