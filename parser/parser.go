package parser

import (
	"fmt"
	"slices"

	"go.creack.net/bistro-matic/ast"
	"go.creack.net/bistro-matic/lexer"
)

type parser struct {
	lex *lexer.Lexer

	curToken lexer.Token

	stmtLookupTable         lookupTable[stmtHandler]
	nudLookupTable          lookupTable[nudHandler]
	ledLookupTable          lookupTable[ledHandler]
	bindingPowerLookupTable lookupTable[bindingPower]
}

func newParser(lex *lexer.Lexer) *parser {
	p := &parser{
		lex: lex,

		stmtLookupTable:         lookupTable[stmtHandler]{},
		nudLookupTable:          lookupTable[nudHandler]{},
		ledLookupTable:          lookupTable[ledHandler]{},
		bindingPowerLookupTable: lookupTable[bindingPower]{},
	}
	p.createTokenLookups()
	return p
}

func Parse(lex *lexer.Lexer) ast.BlockStmt {
	var stms []ast.Stmt

	p := newParser(lex)
	for p.curToken.Type != lexer.TokEOF {
		p.nextToken()

		stms = append(stms, parseStmt(p))
	}

	return ast.BlockStmt{
		Stmts: stms,
	}
}

func (p *parser) nextToken() {
	p.curToken = p.lex.NextToken()
}

func (p *parser) expect(kinds ...lexer.TokenType) {
	if slices.Contains(kinds, p.curToken.Type) {
		p.nextToken()
		return
	}
	panic(fmt.Sprintf("expected token %v, got %s", kinds, p.curToken))
}
