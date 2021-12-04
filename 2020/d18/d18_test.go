package main

import "testing"

func TestInfixToPostfix(t *testing.T) {
	expr := []Token{
		Token{kind: OPERAND, iValue: 1},
		Token{kind: OPERATOR, sValue: "+"},
		Token{kind: OPERAND, iValue: 2},
		Token{kind: OPERATOR, sValue: "*"},
		Token{kind: OPERAND, iValue: 3},
		Token{kind: OPERATOR, sValue: "+"},
		Token{kind: OPERAND, iValue: 4},
		Token{kind: OPERATOR, sValue: "*"},
		Token{kind: OPERAND, iValue: 5},
		Token{kind: OPERATOR, sValue: "+"},
		Token{kind: OPERAND, iValue: 6},
	}

	postfix := infixToPostfix(expr)
	got := eval(postfix)
	want := 71

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// TODO: figure out how to verify that the stack is empty after evaluating...
func TestEval(t *testing.T) {
	tests := []struct {
		expr TokenStream
		want int
	}{
		{expr: []Token{
			Token{kind: OPERAND, iValue: 5},
			Token{kind: OPERAND, iValue: 6},
			Token{kind: OPERATOR, sValue: "+"},
			Token{kind: OPERAND, iValue: 3},
			Token{kind: OPERATOR, sValue: "*"},
			Token{kind: OPERAND, iValue: 4},
			Token{kind: OPERATOR, sValue: "*"},
		}, want: 132},
		{expr: []Token{
			Token{kind: OPERAND, iValue: 4},
			Token{kind: OPERAND, iValue: 5},
			Token{kind: OPERAND, iValue: 6},
			Token{kind: OPERAND, iValue: 7},
			Token{kind: OPERAND, iValue: 8},
			Token{kind: OPERATOR, sValue: "+"},
			Token{kind: OPERATOR, sValue: "*"},
			Token{kind: OPERATOR, sValue: "+"},
			Token{kind: OPERATOR, sValue: "*"},
		}, want: 380},
	}

	for _, tt := range tests {
		got := eval(tt.expr)

		if got != tt.want {
			t.Errorf("got %d, want %d", got, tt.want)
		}
	}
}
