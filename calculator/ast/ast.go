package ast

import (
	"errors"
	"fmt"
	"github.com/eliquious/lexer"
	"math/big"
	"strconv"
	"time"
)

// ExpressionType identifies various expressions
type ExpressionType int

const (
	VariableDeclarationType ExpressionType = iota
	ScopedVariableDeclarationType
	ConstantDeclarationType
	FunctionDeclarationType
	StructDeclarationType
	UnitDeclarationType
	AttributeDeclarationType
	ArrayDeclarationType
	EnumDeclarationType

	ImportExpressionType
	ConversionExpressionType
	IfExpressionType
	IfElseExpressionType
	ElseExpressionType
	CallFunctionExpressionType
	FilterExpressionType
	ForExpressionType

	AssignmentExpressionType
	BinaryExpressionType
	UnaryExpressionType

	IntegerLiteralType
	DecimalLiteralType
	StringLiteralType
	DurationLiteralType
	TimestampLiteralType
	BooleanLiteralType
	StructLiteralType
	ConversionLiteralType
	ArrayLiteralType
)

// Expression represents AST expressions
type Expression interface {
	Type() ExpressionType
	String() string
}

// IsLiteral returns true for literal expressions
func IsLiteral(expr Expression) bool {
	switch expr.Type() {
	case IntegerLiteralType:
		return true
	case DecimalLiteralType:
		return true
	case BooleanLiteralType:
		return true
	case StringLiteralType:
		return true
	case DurationLiteralType:
		return true
	default:
		return false
	}
}

// IsUnaryOperator returns true for unary operators
func IsUnaryOperator(tok lexer.Token) bool {
	if tok == lexer.PLUSPLUS || tok == lexer.MINUSMINUS {
		return true
	}
	return false
}

// IsBinaryOperator returns true for binary operators
func IsBinaryOperator(tok lexer.Token) bool {
	if !tok.IsOperator() {
		return false
	} else if IsUnaryOperator(tok) {
		return false
	}
	return true
}

// IntegerLiteral represents literal integers
type IntegerLiteral struct {
	Value *big.Int
}

func (e IntegerLiteral) Type() ExpressionType { return IntegerLiteralType }
func (e IntegerLiteral) String() string       { return e.Value.String() }

func (e IntegerLiteral) Add(expr Expression) (Expression, error) {
	switch expr.Type() {
	case IntegerLiteralType:
		i := new(big.Int)
		return &IntegerLiteral{i.Add(e.Value, expr.(*IntegerLiteral).Value)}, nil
	case DecimalLiteralType:
		f := new(big.Float).SetInt(e.Value)
		return &DecimalLiteral{f.Add(f, expr.(*DecimalLiteral).Value)}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Integer addition of type '%T' unsupported", expr.Type()))
	}
}

// DecimalLiteral represents literal decimals
type DecimalLiteral struct {
	Value *big.Float
}

func (e DecimalLiteral) Type() ExpressionType { return DecimalLiteralType }
func (e DecimalLiteral) String() string       { return e.Value.Text('E', 16) }

func (e DecimalLiteral) Add(expr Expression) (Expression, error) {
	switch expr.Type() {
	case IntegerLiteralType:
		f := new(big.Float).SetInt(expr.(*IntegerLiteral).Value)
		return &DecimalLiteral{f.Add(f, e.Value)}, nil
	case DecimalLiteralType:
		f := new(big.Float)
		return &DecimalLiteral{f.Add(e.Value, expr.(*DecimalLiteral).Value)}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Decimal addition of type '%T' unsupported", expr.Type()))
	}
}

// BooleanLiteral represents literal booleans
type BooleanLiteral struct {
	Value bool
}

func (e BooleanLiteral) Type() ExpressionType { return BooleanLiteralType }
func (e BooleanLiteral) String() string       { return strconv.FormatBool(e.Value) }

// StringLiteral represents literal strings
type StringLiteral struct {
	Value string
}

func (e StringLiteral) Type() ExpressionType { return StringLiteralType }
func (e StringLiteral) String() string       { return strconv.Quote(e.Value) }

// DurationLiteral represents literal durations
type DurationLiteral struct {
	Value time.Duration
}

func (e DurationLiteral) Type() ExpressionType { return DurationLiteralType }
func (e DurationLiteral) String() string       { return e.Value.String() }
