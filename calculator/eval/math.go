package eval

import (
	"errors"
	"fmt"
	"github.com/eliquious/aechbar/calculator/ast"
	"github.com/eliquious/lexer"
	// "math/big"
)

func evalPlusExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	if expr.Op != lexer.PLUS {
		return nil, errors.New("Expected PLUS operand")
	}

	if add, ok := expr.LExpr.(ast.AddExpression); ok {
		return add.Add(expr.RExpr)
	}
	return nil, errors.New(fmt.Sprintf("PLUS operand not supported for %T and %T", expr.LExpr, expr.RExpr))

	// if lh.Type() == IntegerLiteralType && rh.Type() == IntegerLiteralType {
	// 	i := new(big.Int)
	// 	return &IntegerLiteral{i.Add(lh.(*IntegerLiteral).Value, rh.(*IntegerLiteral).Value)}, nil
	// } else if lh.Type() == IntegerLiteralType && rh.Type() == DecimalLiteralType {
	// 	f := new(big.Float).SetInt(lh.(*IntegerLiteral).Value)
	// 	return &DecimalLiteral{f.Add(f, rh.(*DecimalLiteral).Value)}, nil
	// } else if lh.Type() == DecimalLiteralType && rh.Type() == IntegerLiteralType {
	// 	f := new(big.Float).SetInt(rh.(*IntegerLiteral).Value)
	// 	return &DecimalLiteral{f.Add(f, lh.(*DecimalLiteral).Value)}, nil
	// } else if lh.Type() == DecimalLiteralType && rh.Type() == DecimalLiteralType {
	// 	f := new(big.Float)
	// 	return &DecimalLiteral{f.Add(lh.(*DecimalLiteral).Value, rh.(*DecimalLiteral).Value)}, nil
	// } else if lh.Type() == DurationLiteralType && rh.Type() == DurationLiteralType {
	// 	return &DurationLiteral{lh.(*DurationLiteral).Value + rh.(*DurationLiteral).Value}, nil
	// } else if lh.Type() == StringLiteralType && rh.Type() == StringLiteralType {
	// 	return &StringLiteral{lh.(*StringLiteral).Value + rh.(*StringLiteral).Value}, nil
	// } else {
	// 	return nil, errors.New(fmt.Sprintf("PLUS operand not supported for %T and %T", lh, rh))
	// }
}

func evalMinusExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	if expr.Op != lexer.MINUS {
		return nil, errors.New("Expected MINUS operand")
	}

	if e, ok := expr.LExpr.(ast.SubExpression); ok {
		return e.Sub(expr.RExpr)
	}
	return nil, errors.New(fmt.Sprintf("MINUS operand not supported for %T and %T", expr.LExpr, expr.RExpr))
}

func evalMultExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	if expr.Op != lexer.MUL {
		return nil, errors.New("Expected MUL operand")
	}

	if e, ok := expr.LExpr.(ast.MultExpression); ok {
		return e.Mult(expr.RExpr)
	}
	return nil, errors.New(fmt.Sprintf("MUL operand not supported for %T and %T", expr.LExpr, expr.RExpr))
}

func evalDivExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	if expr.Op != lexer.DIV {
		return nil, errors.New("Expected DIV operand")
	}

	if e, ok := expr.LExpr.(ast.DivExpression); ok {
		return e.Div(expr.RExpr)
	}
	return nil, errors.New(fmt.Sprintf("DIV operand not supported for %T and %T", expr.LExpr, expr.RExpr))
}

func evalPowExpression(expr *ast.BinaryExpression) (ast.Expression, error) {
	if expr.Op != lexer.POW {
		return nil, errors.New("Expected POW operand")
	}

	if e, ok := expr.LExpr.(ast.PowExpression); ok {
		return e.Pow(expr.RExpr)
	}
	return nil, errors.New(fmt.Sprintf("Pow operand not supported for %T and %T", expr.LExpr, expr.RExpr))
}
