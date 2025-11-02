package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-java/bindings/go"
)

func Java() Language {
	return Language{
		Name:       "Java",
		Extensions: []string{".java"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "([(line_comment) (block_comment)] @comment)",
	}
}
