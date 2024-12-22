package calculation

import "errors"

var (
	ErrInvalidExpression = errors.New("invalid expression")
	ErrDivisionByZero    = errors.New("division by zero")
	ErrIncorrectOperator = errors.New("incorrect operator")
	ErrEmptyExpression   = errors.New("empty expression")
)
