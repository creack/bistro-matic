package lexer

import "testing"

func TestTokenTypeStrings(t *testing.T) {
	if len(tokenTypeStrings) != int(FinalToken) {
		t.Errorf("Expected %d token types, got %d", FinalToken, len(tokenTypeStrings))
	}
}

func TestLexer(t *testing.T) {
	input := "1 + 2 - 3 * 4 / 5 % 6 (7)"
	lexer, err := New(input, DefaultBase, DefaultOperators)
	if err != nil {
		t.Fatalf("Failed to create lexer: %s.", err)
	}

	expectedTokens := []Token{
		{TokNumber, "1"},
		{TokPlus, "+"},
		{TokNumber, "2"},
		{TokMinus, "-"},
		{TokNumber, "3"},
		{TokMultiply, "*"},
		{TokNumber, "4"},
		{TokDivide, "/"},
		{TokNumber, "5"},
		{TokModulo, "%"},
		{TokNumber, "6"},
		{TokParenLeft, "("},
		{TokNumber, "7"},
		{TokParenRight, ")"},
	}

	for _, expected := range expectedTokens {
		token := lexer.NextToken()
		if token != expected {
			t.Errorf("Expected token %v, got %v", expected, token)
		}
	}
}
