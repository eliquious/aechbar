package eval

import (
	"errors"
	"github.com/eliquious/aechbar/calculator/ast"
	"github.com/eliquious/lexer"
	"math/big"
)

func Evaluate(expr ast.Expression) (string, error) {
	exp, err := evalExpression(expr)
	if err != nil {
		return "", err
	}
	return exp.String(), nil
}

func evalExpression(expr ast.Expression) (ast.Expression, error) {
	switch expr.Type() {
	case ast.IntegerLiteralType, ast.DecimalLiteralType,
		ast.StringLiteralType, ast.DurationLiteralType,
		ast.TimestampLiteralType, ast.BooleanLiteralType:
		return expr, nil
	case ast.UnaryExpressionType:
		return evalUnaryExpression(expr.(*ast.UnaryExpression))
	case ast.BinaryExpressionType:
		return evalBinaryExpression(expr.(*ast.BinaryExpression))
	default:
		return nil, errors.New("Unsupported expression")
	}
}

func evalUnaryExpression(expr *ast.UnaryExpression) (ast.Expression, error) {
	switch expr.Expr.Type() {
	case ast.IntegerLiteralType:
		return evalUnaryIntegerExpression(expr.Op, expr.Expr.(*ast.IntegerLiteral))
	case ast.DecimalLiteralType:
		return evalUnaryDecimalExpression(expr.Op, expr.Expr.(*ast.DecimalLiteral))
	default:
		return nil, errors.New("Unsupported unary expression")
	}
}

func evalUnaryIntegerExpression(op lexer.Token, expr *ast.IntegerLiteral) (ast.Expression, error) {
	if op == lexer.MINUSMINUS {
		return &ast.IntegerLiteral{expr.Value.Add(expr.Value, big.NewInt(-1))}, nil
	} else if op == lexer.PLUSPLUS {
		return &ast.IntegerLiteral{expr.Value.Add(expr.Value, big.NewInt(1))}, nil
	}
	return nil, errors.New("Unsupported integer unary expression")
}

func evalUnaryDecimalExpression(op lexer.Token, expr *ast.DecimalLiteral) (ast.Expression, error) {
	if op == lexer.MINUSMINUS {
		return &ast.DecimalLiteral{expr.Value.Add(expr.Value, big.NewFloat(-1))}, nil
	} else if op == lexer.PLUSPLUS {
		return &ast.DecimalLiteral{expr.Value.Add(expr.Value, big.NewFloat(1))}, nil
	}
	return nil, errors.New("Unsupported decimal unary expression")
}

func evalBinaryExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	// Reduce the binary expression to it's lowest parts
	exp, err := reduceBinaryExpression(expr)
	if err != nil {
		return nil, err
	}

	switch expr.Op {
	case lexer.PLUS, lexer.MINUS, lexer.MUL, lexer.DIV, lexer.POW:
		return evalBinaryMathExpression(exp)
	case lexer.AMPERSAND, lexer.XOR, lexer.PIPE, lexer.LSHIFT, lexer.RSHIFT:
		return evalBinaryBitwiseExpression(exp)
	case lexer.AND, lexer.OR, lexer.EQEQ, lexer.NEQ, lexer.LT, lexer.LTE, lexer.GT, lexer.GTE:
		return evalBinaryBooleanExpression(exp)
	default:
		return nil, errors.New("Unsupported binary expression")
	}
}

func evalBinaryMathExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	switch expr.Op {
	case lexer.PLUS:
		return evalPlusExpression(expr)
	case lexer.MINUS:
		return evalMinusExpression(expr)
	case lexer.MUL:
		return evalMultExpression(expr)
	case lexer.DIV:
		return evalDivExpression(expr)
	case lexer.POW:
		return evalPowExpression(expr)
	default:
		return nil, errors.New("Unsupported binary expression")
	}
}

func evalBinaryBooleanExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	switch expr.Op {
	case lexer.AND:
	case lexer.OR:
	case lexer.EQEQ:
	case lexer.NEQ:
	case lexer.LT:
	case lexer.LTE:
	case lexer.GT:
	case lexer.GTE:
	default:
		return nil, errors.New("Unsupported boolean expression")
	}
	return nil, errors.New("Unsupported boolean expression")
}

func evalBinaryBitwiseExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	switch expr.Op {
	case lexer.AMPERSAND:
	case lexer.XOR:
	case lexer.PIPE:
	case lexer.LSHIFT:
	case lexer.RSHIFT:
	default:
		return nil, errors.New("Unsupported boolean expression")
	}
	return nil, errors.New("Unsupported boolean expression")
}

func reduceBinaryExpression(expr *ast.BinaryExpression) (*ast.BinaryExpression, error) {

	// Eval left hand side
	lh, err := evalExpression(expr.LExpr)
	if err != nil {
		return nil, err
	}

	// Eval right hand side
	rh, err := evalExpression(expr.RExpr)
	if err != nil {
		return nil, err
	}
	return &ast.BinaryExpression{expr.Op, lh, rh}, nil
}
