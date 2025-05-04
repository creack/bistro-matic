package lexer

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type Lexer struct {
	input string

	pos   int // Current position in input (points to current char).
	start int // Start position of current token.

	atEOF bool

	curToken Token

	Base           string
	operatorTokens map[rune]TokenType
}

func New(input, base, operators string) (*Lexer, error) {
	if len(operators) != 7 {
		return nil, fmt.Errorf("operators string must be 7 characters long")
	}
	if len(base) < 2 {
		return nil, fmt.Errorf("base string must be at least 2 characters long")
	}
	baseSet := map[rune]struct{}{}
	for _, elem := range base {
		baseSet[elem] = struct{}{}
	}
	if len(baseSet) != utf8.RuneCountInString(base) {
		return nil, fmt.Errorf("base string must contain unique characters")
	}
	operatorTokens := map[rune]TokenType{}
	for i, elem := range operators {
		operatorTokens[elem] = OpTable[i]
		if _, ok := baseSet[elem]; ok {
			return nil, fmt.Errorf("base and operator strings must not contain the same characters")
		}
	}
	if len(operatorTokens) != utf8.RuneCountInString(operators) {
		return nil, fmt.Errorf("operators string must contain unique characters")
	}
	return &Lexer{
		input: input,
		pos:   0,

		Base:           base,
		operatorTokens: operatorTokens,
	}, nil
}

func (l *Lexer) NextToken() Token {
	l.curToken = Token{Type: TokEOF, Value: ""}
	state := lexText
	for {
		state = state(l)
		if state == nil {
			// fmt.Printf("LEXER: %s\n", l.curToken)
			return l.curToken
		}
	}

}

func (l *Lexer) next() rune {
	if l.atEOF || l.pos >= len(l.input) {
		l.atEOF = true
		return 0
	}
	r, n := utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += n
	return r
}

func (l *Lexer) backup() {
	if !l.atEOF && l.pos > 0 {
		_, n := utf8.DecodeLastRuneInString(l.input[:l.pos])
		l.pos -= n
	}
}

func (l *Lexer) peek() rune {
	ch := l.next()
	l.backup()
	return ch
}

func (l *Lexer) ignore() {
	l.start = l.pos
}

func (l *Lexer) emit(t TokenType) stateFn {
	l.curToken = Token{Type: t, Value: l.input[l.start:l.pos]}
	l.start = l.pos
	return nil
}

func (l *Lexer) accept(valid string) bool {
	if strings.ContainsRune(valid, l.next()) {
		return true
	}
	l.backup()
	return false
}

func (l *Lexer) acceptRun(valid string) bool {
	accepted := false
	for strings.ContainsRune(valid, l.next()) {
		accepted = true
	}
	l.backup()
	return accepted
}

func (l *Lexer) errorf(format string, args ...any) stateFn {
	l.curToken = Token{Type: TokError, Value: fmt.Sprintf(syntaxErrorMsg+format, args...)}
	l.start = 0
	l.pos = 0
	l.input = l.input[:0]
	l.atEOF = true
	return nil
}

func (l *Lexer) isOperatorToken(r rune) bool {
	_, ok := l.operatorTokens[r]
	return ok
}
