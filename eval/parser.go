package main

import (
	"strings"

	"errors"
	"github.com/eliquious/lexer"
	"io"
)

var EOL = errors.New("End of Line")
var EOF = errors.New("End of Input")

// Parser represents an InfluxQL parser.
type Parser struct {
	s *lexer.TokenBuffer
}

// NewParser returns a new instance of Parser.
func NewParser(r io.Reader) *Parser {
	return &Parser{s: lexer.NewTokenBuffer(r)}
}

func ParseExpression(s string) (Expression, error) {
	return NewParser(strings.NewReader(s)).ParseExpression()
}

// ParseExpression parses a string and returns a Expression AST object.
func (p *Parser) ParseExpression() (Expression, error) {

	// Inspect the first token.
	tok, pos, lit := p.scanIgnoreWhitespace()
	switch tok {
	// case VAR, LET, CONST:
	// 	return p.parseAssignment(tok)
	// case FUNC:
	// 	return p.parseFunctionDeclaration()
	// case STRUCT:
	// 	return p.parseStructDeclaration()
	// case UNIT:
	// 	return p.parseUnitDeclaration()
	// case IF:
	// 	return p.parseIfExpression()
	// case ELSEIF:
	// 	return p.parseElseIfExpression()
	// case ELSE:
	// 	return p.parseElseExpression()
	// case IMPORT:
	// 	return p.parseImportExpression()
	// case FOR:
	// 	return p.parseForExpression()
	case lexer.INTEGER, lexer.DECIMAL, lexer.STRING, lexer.TRUE, lexer.FALSE, lexer.DURATION:
		return p.parseLiteralExpression(tok, pos, lit)
	case lexer.PLUS, lexer.MINUS:
		tok2, _, lit2 := p.scanIgnoreWhitespace()
		if tok2 == lexer.INTEGER {
			return p.parseLiteralInteger(tok2, pos, lit+lit2)
		} else if tok2 == lexer.DECIMAL {
			return p.parseLiteralDecimal(tok2, pos, lit+lit2)
		}
		p.unscan()
		return nil, tokenError("Invalid input", tok, pos, lit)

	// case lexer.IDENT:
	// 	return p.parseIdentExpression()
	// case lexer.LPAREN:
	// 	return p.parseParenExpression()
	// case lexer.LBRACKET:
	// 	return p.parseArrayExpression()
	case lexer.SEMICOLON:
		return nil, EOL
	case lexer.EOF:
		return nil, EOF
	default:
		return nil, tokenError("Unrecognized input error", tok, pos, lit)
		// return nil, newParseError(tokstr(tok, lit), []string{"USE", "CREATE", "SHOW", "DROP"}, pos)
	}
}

// scan returns the next token from the underlying scanner.
func (p *Parser) scan() (tok lexer.Token, pos lexer.Pos, lit string) { return p.s.Scan() }

// unscan pushes the previously read token back onto the buffer.
func (p *Parser) unscan() { p.s.Unscan() }

// // peekRune returns the next rune that would be read by the scanner.
// func (p *Parser) peekRune() rune { return p.s.s.Peek() }

// scanIgnoreWhitespace scans the next non-whitespace token.
func (p *Parser) scanIgnoreWhitespace() (tok lexer.Token, pos lexer.Pos, lit string) {
	tok, pos, lit = p.scan()
	if tok == lexer.WS {
		tok, pos, lit = p.scan()
	}
	return
}
