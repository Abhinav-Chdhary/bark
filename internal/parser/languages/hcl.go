package languages

import (
	tree_sitter "github.com/tree-sitter-grammars/tree-sitter-hcl/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

func HCL() Language {
	return Language{
		Name:       "HCL",
		Extensions: []string{".hcl", ".tf", ".tfvars"},
		Parser:     sitter.NewLanguage(tree_sitter.Language()),
		Query:      "((comment) @comment)",
	}
}
