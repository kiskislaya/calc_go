package calculation

import (
	"strconv"
	"strings"
)

func isDigit(char rune) bool {
	if char >= '0' && char <= '9' {
		return true
	}
	return false
}

func isOp(char rune) bool {
	if char == '+' || char == '-' || char == '*' || char == '/' {
		return true
	}
	return false
}

func parseNumber(s string, i int) (float64, int) {
	num := ""
	for i < len(s) {
		if (s[i] >= '0' && s[i] <= '9') || s[i] == '.' {
			num += string(s[i])
		} else {
			break
		}
		i++

	}
	res, err := strconv.ParseFloat(num, 64)
	if err != nil {
		return 0, -1
	}
	return res, i

}

func priorityOp(op rune) int {
	switch op {
	case '+':
		return 1
	case '-':
		return 1
	case '*':
		return 2
	case '/':
		return 2
	default:
		return -1
	}
}

func calculate(numStack *[]float64, opStack *[]rune) error {
	if len(*numStack) < 2 || len(*opStack) == 0 {
		return ErrIncorrectOperator
	}

	var res float64

	num1 := (*numStack)[len(*numStack)-2]
	num2 := (*numStack)[len(*numStack)-1]
	op := (*opStack)[len(*opStack)-1]

	*numStack = (*numStack)[:len(*numStack)-2]
	*opStack = (*opStack)[:len(*opStack)-1]

	switch op {
	case '+':
		res = num1 + num2
	case '-':
		res = num1 - num2
	case '*':
		res = num1 * num2
	case '/':
		if num2 == 0 {
			return ErrDivisionByZero
		}
		res = num1 / num2
	default:
		return ErrIncorrectOperator
	}

	*numStack = append(*numStack, res)
	return nil
}

func Calc(expression string) (float64, error) {
	if len(expression) == 0 {
		return 0, ErrEmptyExpression
	}
	expression = strings.ReplaceAll(expression, " ", "")

	numStack := make([]float64, 0)
	opStack := make([]rune, 0)

	for i := 0; i < len(expression); i++ {
		char := rune(expression[i])
		if isDigit(char) {
			num, j := parseNumber(expression, i)
			numStack = append(numStack, num)
			i = j - 1
		} else if isOp(char) {
			for len(opStack) > 0 && priorityOp(opStack[len(opStack)-1]) >= priorityOp(char) {
				if err := calculate(&numStack, &opStack); err != nil {
					return 0, err
				}
			}
			opStack = append(opStack, char)
		} else if char == '(' {
			opStack = append(opStack, char)
		} else if char == ')' {
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				if err := calculate(&numStack, &opStack); err != nil {
					return 0, err
				}
			}
			if len(opStack) == 0 {
				return 0, ErrInvalidExpression
			}
			opStack = opStack[:len(opStack)-1]
		} else {
			return 0, ErrInvalidExpression
		}
	}

	for len(opStack) > 0 {
		if err := calculate(&numStack, &opStack); err != nil {
			return 0, err
		}
	}

	if len(numStack) > 1 {
		return 0, ErrInvalidExpression
	}

	return numStack[0], nil
}
