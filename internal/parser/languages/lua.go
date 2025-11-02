package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-lua/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func Lua() Language {
	return Language{
		Name:       "Lua",
		Extensions: []string{".lua"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
