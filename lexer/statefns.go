package lexer

import (
	"strings"
)

type stateFn func(*Lexer) stateFn

func lexText(l *Lexer) stateFn {
	whitespaces := ""
	if !strings.ContainsRune(l.Base, '\t') {
		whitespaces += "\t"
	}
	if !strings.ContainsRune(l.Base, ' ') {
		whitespaces += " "
	}
	if !strings.ContainsRune(l.Base, '\r') {
		whitespaces += "\r"
	}
	if !strings.ContainsRune(l.Base, '\n') {
		whitespaces += "\n"
	}
	l.acceptRun(whitespaces)
	l.ignore()
	switch r := l.peek(); {
	case r == 0:
		return l.emit(TokEOF)
	case strings.ContainsRune(l.Base, r):
		return lexNumber
	case l.isOperatorToken(r):
		l.next()
		return l.emit(l.operatorTokens[r])
	default:
		return l.errorf("unexpected character: %q", r)
	}
}

func lexNumber(l *Lexer) stateFn {
	l.acceptRun(l.Base)
	l.accept(".")
	l.acceptRun(l.Base)
	return l.emit(TokNumber)
}
