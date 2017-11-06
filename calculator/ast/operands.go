package ast

type AddExpression interface {
	Add(Expression) (Expression, error)
}

type SubExpression interface {
	Sub(Expression) (Expression, error)
}

type MultExpression interface {
	Mult(Expression) (Expression, error)
}

type DivExpression interface {
	Div(Expression) (Expression, error)
}

type PowExpression interface {
	Pow(Expression) (Expression, error)
}

type AmpersandExpression interface {
	Ampersand(Expression) (Expression, error)
}

type XorExpression interface {
	Xor(Expression) (Expression, error)
}

type PipeExpression interface {
	Pipe(Expression) (Expression, error)
}

type LShiftExpression interface {
	LShift(Expression) (Expression, error)
}

type RShiftExpression interface {
	RShift(Expression) (Expression, error)
}

type AndExpression interface {
	And(Expression) (Expression, error)
}

type OrExpression interface {
	Or(Expression) (Expression, error)
}

type EqualExpression interface {
	Equal(Expression) (Expression, error)
}

type NotEqualExpression interface {
	NotEqual(Expression) (Expression, error)
}

type LessThanExpression interface {
	LessThan(Expression) (Expression, error)
}

type LessThanEqualToExpression interface {
	LessThanOrEqualTo(Expression) (Expression, error)
}

type GreaterThanExpression interface {
	GreaterThan(Expression) (Expression, error)
}

type GreaterThanEqualToExpression interface {
	GreaterThanOrEqualTo(Expression) (Expression, error)
}
