package ast

import "go.creack.net/bistro-matic/lexer"

type Expr interface {
	expr()
}

type NumberExpr struct {
	Value int
}

func (NumberExpr) expr() {}

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (BinaryExpr) expr() {}

type PrefixExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (PrefixExpr) expr() {}

type Stmt interface {
	stmt()
}

type BlockStmt struct {
	Stmts []Stmt
}

func (BlockStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (ExpressionStmt) stmt() {}
