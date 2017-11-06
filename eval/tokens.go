package main

import "github.com/eliquious/lexer"

func init() {
	lexer.LoadTokenMap(keywords)
}

// Token enums
// Built-in Types
const (
	startTypes lexer.Token = iota + 1000
	STRING_TYPE
	INT_TYPE
	FLOAT_TYPE
	TIMESTAMP_TYPE
	DURATION_TYPE
	BOOLEAN_TYPE
	endTypes

	// Keywords
	startKeywords

	// Var types
	FUNC
	VAR
	LET
	CONST
	STRUCT
	ENUM

	// Units
	UNIT
	CONVERSION
	TO

	//
	IF
	ELSE
	ELSEIF
	FOR
	FILTER
	IMPORT

	endKeywords

	// Built-in functions
	startFunctions

	APPEND
	LEN
	POP
	PUSH
	SQRT
	LOG

	endFunctions
)

var keywords = map[lexer.Token]string{

	// Base Types
	STRING_TYPE:    "string",
	INT_TYPE:       "int",
	FLOAT_TYPE:     "float",
	TIMESTAMP_TYPE: "timestamp",
	DURATION_TYPE:  "duration",
	BOOLEAN_TYPE:   "boolean",

	// Var types
	FUNC:   "FUNC",
	VAR:    "VAR",
	LET:    "LET",
	CONST:  "CONST",
	STRUCT: "STRUCT",

	// Units
	UNIT:       "UNIT",
	CONVERSION: "CONVERSION",
	TO:         "TO",

	// Expressions
	IF:     "IF",
	ELSE:   "ELSE",
	FOR:    "FOR",
	FILTER: "FILTER",
	ENUM:   "ENUM",

	// Functions
	APPEND: "APPEND",
	LEN:    "LEN",
	POP:    "POP",
	PUSH:   "PUSH",
	SQRT:   "SQRT",
	LOG:    "LOG",
}

// tokstr returns a literal if provided, otherwise returns the token string.
func tokstr(tok lexer.Token, lit string) string {
	if lit != "" && tok != lexer.WS {
		return lit
	}
	return tok.String()
}
