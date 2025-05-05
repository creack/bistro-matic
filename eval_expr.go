package main

import (
	"fmt"

	"go.creack.net/bistro-matic/ast"
	"go.creack.net/bistro-matic/lexer"
	"go.creack.net/bistro-matic/parser"
)

func evalExpr(base, ops, expr string, _ int) (string, error) {
	// Create a new lexer with the given base and operators.
	lex, err := lexer.New(expr, base, ops)
	if err != nil {
		return "", fmt.Errorf("new lexer: %w", err)
	}
	// Parse the expression into an AST.
	tree := parser.Parse(lex)
	if len(tree.Stmts) != 1 {
		return "", fmt.Errorf("expected one statement, got %d", len(tree.Stmts))
	}
	exprStmt, ok := tree.Stmts[0].(ast.ExpressionStmt)
	if !ok {
		return "", fmt.Errorf("expected expression statement, got %T", tree.Stmts[0])
	}

	// Evaluate the AST.
	result := Evaluate(exprStmt.Expression)
	return parser.PutNumberBase(int(result), base)
}

// Evaluate evaluates an AST node and returns its value.
func Evaluate(node ast.Expr) int {
	switch n := node.(type) {
	case ast.NumberExpr:
		return n.Value
	case ast.BinaryExpr:
		return evaluateBinaryOpNode(n)
	case ast.PrefixExpr:
		return evaluatePrefixNode(n)
	default:
		panic(fmt.Errorf("unknown node type %T", n))
	}
}

// evaluateBinaryOpNode evaluates a binary operation node
func evaluateBinaryOpNode(node ast.BinaryExpr) int {
	left := Evaluate(node.Left)
	right := Evaluate(node.Right)

	switch node.Operator.Type {
	case lexer.TokPlus:
		return left + right
	case lexer.TokMinus:
		return left - right
	case lexer.TokMultiply:
		return left * right
	case lexer.TokDivide:
		if right == 0 {
			panic("division by zero")
		}
		return left / right
	case lexer.TokModulo:
		if right == 0 {
			panic("modulo by zero")
		}
		return left % right
	default:
		panic(fmt.Sprintf("unknown operator: %s", node.Operator))
	}
}

// evaluatePrefixNode evaluates a unary operation node
func evaluatePrefixNode(node ast.PrefixExpr) int {
	operand := Evaluate(node.Right)

	switch node.Operator.Type {
	case lexer.TokMinus:
		return -operand
	case lexer.TokPlus:
		return operand
	default:
		panic(fmt.Errorf("unknown prefix operator: %s", node.Operator.Type))
	}
}
