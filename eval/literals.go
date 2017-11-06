package main

import (
	"fmt"
	"github.com/eliquious/lexer"
	"math/big"
	"time"
)

func (p *Parser) parseLiteralExpression(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	expr, err := p.parseLiteral(tok, pos, lit)
	if err != nil {
		return nil, err
	}

	// Scan for operators
	tok, pos, lit = p.scanIgnoreWhitespace()
	if IsUnaryOperator(tok) {
		return &UnaryExpression{tok, expr}, nil
	} else if IsBinaryOperator(tok) {
		return p.parseBinaryExpression(expr, tok, pos, lit)
	} else {
		p.unscan()
	}
	return expr, nil
}

func (p *Parser) parseBinaryExpression(lh Expression, op lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	if !IsBinaryOperator(op) {
		return nil, tokenError("Invalid binary operator", op, pos, lit)
	}

	// Process right hand expression
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	// If literal or unary expression, set to right hand side.
	if IsLiteral(expr) || expr.Type() == UnaryExpressionType {
		return &BinaryExpression{Op: op, LExpr: lh, RExpr: expr}, nil
	} else if expr.Type() == BinaryExpressionType {
		return handleBinaryPrecedence(lh, op, expr.(*BinaryExpression))
	} else {
		return nil, fmt.Errorf("Unsupported binary expression: %s", expr.String())
	}
}

// Swap binary expressions based on precedence
func handleBinaryPrecedence(lh Expression, op lexer.Token, rh *BinaryExpression) (Expression, error) {
	if op.Precedence() < rh.Precedence() {
		return &BinaryExpression{op, lh, rh}, nil
	} else if op.Precedence() > rh.Precedence() {
		return &BinaryExpression{rh.Op, &BinaryExpression{op, lh, rh.LExpr}, rh.RExpr}, nil
	}
	return &BinaryExpression{op, lh, rh}, nil
}

func (p *Parser) parseLiteral(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	switch tok {
	case lexer.INTEGER:
		return p.parseLiteralInteger(tok, pos, lit)
	case lexer.DECIMAL:
		return p.parseLiteralDecimal(tok, pos, lit)
	case lexer.STRING:
		return p.parseLiteralString(tok, pos, lit)
	case lexer.FALSE, lexer.TRUE:
		return p.parseLiteralBoolean(tok, pos, lit)
	case lexer.DURATION:
		return p.parseLiteralDuration(tok, pos, lit)
	default:
		return nil, tokenError("Unrecognized literal token", tok, pos, lit)
	}
}

func (p *Parser) parseLiteralInteger(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	i := new(big.Int)
	_, err := fmt.Sscan(lit, i)
	if err != nil {
		return nil, tokenError("Integer literal parse error", tok, pos, lit)
	}
	return &IntegerLiteral{i}, nil
}

func (p *Parser) parseLiteralDecimal(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	f := new(big.Float)
	_, err := fmt.Sscan(lit, f)
	if err != nil {
		return nil, tokenError("Decimal literal parse error", tok, pos, lit)
	}
	return &DecimalLiteral{f}, nil
}

func (p *Parser) parseLiteralBoolean(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	switch tok {
	case lexer.TRUE:
		return &BooleanLiteral{true}, nil
	case lexer.FALSE:
		return &BooleanLiteral{false}, nil
	default:
		return nil, tokenError("Invalid boolean literal", tok, pos, lit)
	}
}

func (p *Parser) parseLiteralString(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	return &StringLiteral{lit}, nil
}

func (p *Parser) parseLiteralDuration(tok lexer.Token, pos lexer.Pos, lit string) (Expression, error) {
	duration, err := time.ParseDuration(lit)
	if err != nil {
		return nil, tokenError("Invalid duration literal", tok, pos, lit)
	}
	return &DurationLiteral{duration}, nil
}
