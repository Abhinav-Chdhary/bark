package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-javascript/bindings/go"
)

func JavaScript() Language {
	return Language{
		Name:       "JavaScript",
		Extensions: []string{".js", ".jsx", ".mjs", ".cjs"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
