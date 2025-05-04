package parser

import (
	"go.creack.net/bistro-matic/ast"
	"go.creack.net/bistro-matic/lexer"
)

type stmtHandler func(*parser) ast.Stmt
type nudHandler func(*parser) ast.Expr
type ledHandler func(*parser, ast.Expr, bindingPower) ast.Expr

type lookupTable[T any] map[lexer.TokenType]T

func (p *parser) led(kind lexer.TokenType, bp bindingPower, fn ledHandler) {
	if _, ok := p.ledLookupTable[kind]; ok {
		panic("duplicate led handler")
	}
	p.ledLookupTable[kind] = fn
	p.bindingPowerLookupTable[kind] = bp
}

func (p *parser) nud(kind lexer.TokenType, fn nudHandler) {
	if _, ok := p.nudLookupTable[kind]; ok {
		panic("duplicate nud handler")
	}
	p.nudLookupTable[kind] = fn
}

func (p *parser) createTokenLookups() {
	// Additional & multiplicative.
	p.led(lexer.TokPlus, bpAdditive, parseBinaryExpr)
	p.led(lexer.TokMinus, bpAdditive, parseBinaryExpr)
	p.led(lexer.TokMultiply, bpMultiplicative, parseBinaryExpr)
	p.led(lexer.TokModulo, bpMultiplicative, parseBinaryExpr)
	p.led(lexer.TokDivide, bpMultiplicative, parseBinaryExpr)

	// Literals & symbols.
	p.nud(lexer.TokNumber, parsePrimaryExpr)
	p.nud(lexer.TokParenLeft, parseGroupingExpr)
	p.nud(lexer.TokPlus, parsePrefixExpr)
	p.nud(lexer.TokMinus, parsePrefixExpr)
}
