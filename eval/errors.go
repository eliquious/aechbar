package main

import (
	"fmt"
	"strings"

	"github.com/eliquious/lexer"
)

type StatusCode int

// Success codes
const (
	OK StatusCode = iota + 2000
)

// Error codes
const (
	InternalServerError StatusCode = iota + 5000
)

var statusCodes = map[StatusCode]string{

	// Success
	OK: "OK",

	// General errors
	InternalServerError: "InternalServerError",
}

// ParseError represents an error that occurred during parsing.
type ParseError struct {
	Message  string
	Found    string
	Expected []string
	Pos      lexer.Pos
}

// newParseError returns a new instance of ParseError.
func newParseError(found string, expected []string, pos lexer.Pos) *ParseError {
	return &ParseError{Found: found, Expected: expected, Pos: pos}
}

// Error returns the string representation of the error.
func (e *ParseError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("%s at line %d, char %d", e.Message, e.Pos.Line+1, e.Pos.Char+1)
	}
	return fmt.Sprintf("found %s, expected %s at line %d, char %d", e.Found, strings.Join(e.Expected, ", "), e.Pos.Line+1, e.Pos.Char+1)
}

func tokenError(message string, tok lexer.Token, pos lexer.Pos, lit string) error {
	return fmt.Errorf("%s : %s at (%d, %d)", message, tokstr(tok, lit), pos.Line+1, pos.Char+1)
}
