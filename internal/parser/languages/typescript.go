package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-typescript/bindings/go"
)

func TypeScript() Language {
	return Language{
		Name:       "TypeScript",
		Extensions: []string{".ts", ".tsx"},
		Parser:     sitter.NewLanguage(tree_sitter.LanguageTSX()),
		Query:      "((comment) @comment)",
	}
}
