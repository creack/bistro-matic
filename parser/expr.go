package parser

import (
	"fmt"

	"go.creack.net/bistro-matic/ast"
	"go.creack.net/bistro-matic/lexer"
)

func parseExpr(p *parser, bp bindingPower) ast.Expr {
	// Parse the primary expression, always start with nud.
	nudFn, exists := p.nudLookupTable[p.curToken.Type]
	if !exists {
		// TODO: Handle errors properly.
		panic(fmt.Errorf("nud handler not found for %s", p.curToken))
	}
	left := nudFn(p)

	// While we have tokens with a higher binding power, parse them using led.
	for p.bindingPowerLookupTable[p.curToken.Type] > bp {
		ledFn, exists := p.ledLookupTable[p.curToken.Type]
		if !exists {
			// TODO: Handle errors properly.
			panic("led handler not found")
		}
		left = ledFn(p, left, p.bindingPowerLookupTable[p.curToken.Type])
	}

	return left
}

func parsePrimaryExpr(p *parser) ast.Expr {
	switch p.curToken.Type {
	case lexer.TokNumber:
		val := p.curToken.Value
		number, err := ParseNumberBase(val, p.lex.Base)
		if err != nil {
			// TODO: Handle errors properly.
			panic("invalid number")
		}
		p.nextToken()
		return ast.NumberExpr{
			Value: number,
		}
	default:
		// TODO: Handle errors properly.
		panic("unexpected token")
	}
}

func parseBinaryExpr(p *parser, left ast.Expr, bp bindingPower) ast.Expr {
	operator := p.curToken
	p.nextToken()
	right := parseExpr(p, bp)

	return ast.BinaryExpr{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func parsePrefixExpr(p *parser) ast.Expr {
	operator := p.curToken
	p.nextToken()
	right := parseExpr(p, bpDefault)

	return ast.PrefixExpr{
		Operator: operator,
		Right:    right,
	}
}

func parseGroupingExpr(p *parser) ast.Expr {
	p.nextToken() // consume '('.
	expr := parseExpr(p, bpDefault)
	p.expect(lexer.TokParenRight) // consumes ')'.
	return expr
}
