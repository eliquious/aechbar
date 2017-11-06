package parser

import (
	"fmt"
	"github.com/eliquious/aechbar/calculator/ast"
	"github.com/eliquious/lexer"
	"math/big"
	"time"
)

func (p *Parser) parseLiteralExpression(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	expr, err := p.parseLiteral(tok, pos, lit)
	if err != nil {
		return nil, err
	}

	// Scan for operators
	tok, pos, lit = p.scanIgnoreWhitespace()
	if ast.IsUnaryOperator(tok) {
		return &ast.UnaryExpression{tok, expr}, nil
	} else if ast.IsBinaryOperator(tok) {
		return p.parseBinaryExpression(expr, tok, pos, lit)
	} else {
		p.unscan()
	}
	return expr, nil
}

func (p *Parser) parseBinaryExpression(lh ast.Expression, op lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	if !ast.IsBinaryOperator(op) {
		return nil, tokenError("Invalid binary operator", op, pos, lit)
	}

	// Process right hand expression
	expr, err := p.ParseExpression()
	if err != nil {
		return nil, err
	}

	// If literal or unary expression, set to right hand side.
	if ast.IsLiteral(expr) || expr.Type() == ast.UnaryExpressionType {
		return &ast.BinaryExpression{Op: op, LExpr: lh, RExpr: expr}, nil
	} else if expr.Type() == ast.BinaryExpressionType {
		return handleBinaryPrecedence(lh, op, expr.(*ast.BinaryExpression))
	} else {
		return nil, fmt.Errorf("Unsupported binary expression: %s", expr.String())
	}
}

// Swap binary expressions based on precedence
func handleBinaryPrecedence(lh ast.Expression, op lexer.Token, rh *ast.BinaryExpression) (ast.Expression, error) {
	if op.Precedence() < rh.Precedence() {
		return &ast.BinaryExpression{op, lh, rh}, nil
	} else if op.Precedence() > rh.Precedence() {
		return &ast.BinaryExpression{rh.Op, &ast.BinaryExpression{op, lh, rh.LExpr}, rh.RExpr}, nil
	}
	return &ast.BinaryExpression{op, lh, rh}, nil
}

func (p *Parser) parseLiteral(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
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

func (p *Parser) parseLiteralInteger(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	i := new(big.Int)
	_, err := fmt.Sscan(lit, i)
	if err != nil {
		return nil, tokenError("Integer literal parse error", tok, pos, lit)
	}
	return &ast.IntegerLiteral{i}, nil
}

func (p *Parser) parseLiteralDecimal(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	f := new(big.Float)
	_, err := fmt.Sscan(lit, f)
	if err != nil {
		return nil, tokenError("Decimal literal parse error", tok, pos, lit)
	}
	return &ast.DecimalLiteral{f}, nil
}

func (p *Parser) parseLiteralBoolean(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	switch tok {
	case lexer.TRUE:
		return &ast.BooleanLiteral{true}, nil
	case lexer.FALSE:
		return &ast.BooleanLiteral{false}, nil
	default:
		return nil, tokenError("Invalid boolean literal", tok, pos, lit)
	}
}

func (p *Parser) parseLiteralString(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	return &ast.StringLiteral{lit}, nil
}

func (p *Parser) parseLiteralDuration(tok lexer.Token, pos lexer.Pos, lit string) (ast.Expression, error) {
	duration, err := time.ParseDuration(lit)
	if err != nil {
		return nil, tokenError("Invalid duration literal", tok, pos, lit)
	}
	return &ast.DurationLiteral{duration}, nil
}
