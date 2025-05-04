package lexer

type TokenType int

const (
	TokError TokenType = iota
	TokEOF

	// Literals.
	TokNumber

	// Operators.
	TokPlus
	TokMinus
	TokMultiply
	TokDivide
	TokModulo

	// Separators.
	TokParenLeft
	TokParenRight
	TokSemicolon
	TokNewline

	FinalToken
)

func (t TokenType) String() string {
	return tokenTypeStrings[t]
}

var tokenTypeStrings = map[TokenType]string{
	TokError: "ERROR",
	TokEOF:   "EOF",

	TokNumber: "NUMBER",

	TokPlus:     "PLUS",
	TokMinus:    "MINUS",
	TokMultiply: "MULTIPLY",
	TokDivide:   "DIVIDE",
	TokModulo:   "MODULO",

	TokParenLeft:  "PAREN_LEFT",
	TokParenRight: "PAREN_RIGHT",
	TokSemicolon:  "SEMICOLON",
	TokNewline:    "NEWLINE",
}
