package parser

import (
	"testing"

	"github.com/kr/pretty"
	"go.creack.net/bistro-matic/lexer"
)

func TestParser(t *testing.T) {
	// Create a lexer with some test input.
	lex, err := lexer.New("10 * 10 + (45 / 10 - -5)", lexer.DefaultBase, lexer.DefaultOperators)
	if err != nil {
		t.Fatalf("Failed to create lexer: %s.", err)
	}

	// Parse the input.
	block := Parse(lex)

	_, _ = pretty.Println(block)
}
