package main

import (
	"errors"
	"github.com/eliquious/lexer"
	"math/big"
)

func Evaluate(expr Expression) (string, error) {
	exp, err := evalExpression(expr)
	if err != nil {
		return "", err
	}
	return exp.String(), nil
}

func evalExpression(expr Expression) (Expression, error) {
	switch expr.Type() {
	case IntegerLiteralType, DecimalLiteralType,
		StringLiteralType, DurationLiteralType,
		TimestampLiteralType, BooleanLiteralType:
		return expr, nil
	case UnaryExpressionType:
		return evalUnaryExpression(expr.(*UnaryExpression))
	case BinaryExpressionType:
		return evalBinaryExpression(expr.(*BinaryExpression))
	default:
		return nil, errors.New("Unsupported expression")
	}
}

func evalUnaryExpression(expr *UnaryExpression) (Expression, error) {
	switch expr.Expr.Type() {
	case IntegerLiteralType:
		return evalUnaryIntegerExpression(expr.Op, expr.Expr.(*IntegerLiteral))
	case DecimalLiteralType:
		return evalUnaryDecimalExpression(expr.Op, expr.Expr.(*DecimalLiteral))
	default:
		return nil, errors.New("Unsupported unary expression")
	}
}

func evalUnaryIntegerExpression(op lexer.Token, expr *IntegerLiteral) (Expression, error) {
	if op == lexer.MINUSMINUS {
		return &IntegerLiteral{expr.Value.Add(expr.Value, big.NewInt(-1))}, nil
	} else if op == lexer.PLUSPLUS {
		return &IntegerLiteral{expr.Value.Add(expr.Value, big.NewInt(1))}, nil
	}
	return nil, errors.New("Unsupported integer unary expression")
}

func evalUnaryDecimalExpression(op lexer.Token, expr *DecimalLiteral) (Expression, error) {
	if op == lexer.MINUSMINUS {
		return &DecimalLiteral{expr.Value.Add(expr.Value, big.NewFloat(-1))}, nil
	} else if op == lexer.PLUSPLUS {
		return &DecimalLiteral{expr.Value.Add(expr.Value, big.NewFloat(1))}, nil
	}
	return nil, errors.New("Unsupported decimal unary expression")
}

func evalBinaryExpression(expr *BinaryExpression) (Expression, error) {
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

func evalBinaryMathExpression(expr *BinaryExpression) (Expression, error) {
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

func evalBinaryBooleanExpression(expr *BinaryExpression) (Expression, error) {
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

func evalBinaryBitwiseExpression(expr *BinaryExpression) (Expression, error) {
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

func reduceBinaryExpression(expr *BinaryExpression) (*BinaryExpression, error) {

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
	return &BinaryExpression{expr.Op, lh, rh}, nil
}
