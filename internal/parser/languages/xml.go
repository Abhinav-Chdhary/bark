package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-xml/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func XML() Language {
	return Language{
		Name:       "XML",
		Extensions: []string{".xml"},
		Parser:     sitter.NewLanguage(tree_sitter.LanguageXML()),
		Query:      "((Comment) @comment)",
	}
}
