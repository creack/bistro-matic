package parser

import (
	"go.creack.net/bistro-matic/ast"
	"go.creack.net/bistro-matic/lexer"
)

func parseStmt(p *parser) ast.Stmt {
	expression := parseExpr(p, bpDefault)
	p.expect(lexer.TokSemicolon, lexer.TokNewline, lexer.TokEOF)

	return ast.ExpressionStmt{
		Expression: expression,
	}
}
