package tree_sitter_motoko_test

import (
	"testing"

	tree_sitter "github.com/smacker/go-tree-sitter"
	"github.com/tree-sitter/tree-sitter-motoko"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_motoko.Language())
	if language == nil {
		t.Errorf("Error loading Motoko grammar")
	}
}
