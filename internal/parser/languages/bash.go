package languages

import (
	sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter "github.com/tree-sitter/tree-sitter-bash/bindings/go"
)

func Bash() Language {
	return Language{
		Name:       "Bash",
		Extensions: []string{".sh", ".bash"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
