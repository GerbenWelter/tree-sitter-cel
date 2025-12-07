package tree_sitter_cel_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_cel "github.com/tree-sitter/tree-sitter-cel/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_cel.Language())
	if language == nil {
		t.Errorf("Error loading Cel grammar")
	}
}
