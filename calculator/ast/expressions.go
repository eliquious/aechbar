package ast

import (
	"fmt"
	"github.com/eliquious/lexer"
)

type UnaryExpression struct {
	Op   lexer.Token
	Expr Expression
}

func (e UnaryExpression) Type() ExpressionType { return UnaryExpressionType }
func (e UnaryExpression) String() string       { return e.Expr.String() + e.Op.String() }

type BinaryExpression struct {
	Op    lexer.Token
	LExpr Expression
	RExpr Expression
}

func (e BinaryExpression) Precedence() int      { return e.Op.Precedence() }
func (e BinaryExpression) Type() ExpressionType { return BinaryExpressionType }
func (e BinaryExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", e.LExpr.String(), e.Op.String(), e.RExpr.String())
}
