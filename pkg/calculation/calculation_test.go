package calculation_test

import (
	"testing"

	"github.com/kiskislaya/calc_go/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
		{
			name:           "complex",
			expression:     "3+(6*2)/3-4",
			expectedResult: 3,
		},
		{
			name:           "nested parentheses",
			expression:     "((2+3)*2)/5",
			expectedResult: 2,
		},
		{
			name:           "multiple operations",
			expression:     "10-2+3*4/2",
			expectedResult: 14,
		},
		{
			name:           "negative result",
			expression:     "2-5",
			expectedResult: -3,
		},
		{
			name:           "decimal numbers",
			expression:     "1.5+2.5",
			expectedResult: 4,
		},
		{
			name:           "division with decimal",
			expression:     "5/2",
			expectedResult: 2.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "simple",
			expression: "1+1*",
		},
		{
			name:       "priority",
			expression: "2+2**2",
		},
		{
			name:       "priority",
			expression: "((2+2-*(2",
		},
		{
			name:       "empty",
			expression: "",
		},
		{
			name:       "division by zero",
			expression: "1/0",
		},
		{
			name:       "invalid character",
			expression: "2+2a",
		},
		{
			name:       "unmatched parentheses",
			expression: "(2+3",
		},
		{
			name:       "multiple operators",
			expression: "2++2",
		},
		{
			name:       "leading operator",
			expression: "+2+2",
		},
		{
			name:       "trailing operator",
			expression: "2+2-",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}
}
