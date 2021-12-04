package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/collections/stack"
)

func readData() (expressions []TokenStream) {
	scanner := bufio.NewScanner(os.Stdin)
	expressions = make([]TokenStream, 0)

	for scanner.Scan() {
		line := scanner.Text()
		ts := make([]Token, 0)

		for _, c := range strings.Split(line, "") {
			var token Token

			switch c {
			case " ":
				continue
			case "(":
				token = Token{kind: LPAREN}
			case ")":
				token = Token{kind: RPAREN}
			case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
				val, _ := strconv.Atoi(c)
				token = Token{kind: OPERAND, iValue: val}
			case "+", "*":
				token = Token{kind: OPERATOR, sValue: c}
			}

			ts = append(ts, token)
		}

		expressions = append(expressions, ts)
	}

	return
}

// Takes expression and converts to postfix
func infixToPostfix1(expression TokenStream) (result TokenStream) {
	result = make(TokenStream, 0)
	s := stack.New()

	for _, t := range expression {
		switch t.kind {
		case OPERAND:
			result = append(result, t)
			if s.Len() > 0 {
				op := s.Peek().(Token)
				if op.kind == OPERATOR {
					result = append(result, op)
					s.Pop()
				}
			}
		case OPERATOR, LPAREN:
			s.Push(t)
		case RPAREN:
			_ = s.Pop().(Token) // this should be an lparen

			for s.Len() > 0 && s.Peek().(Token).kind == OPERATOR {
				op := s.Pop().(Token)
				result = append(result, op)
			}
		}
	}

	for s.Len() > 0 {
		t := s.Pop().(Token)
		result = append(result, t)
	}

	return
}

func precedence(operator string) int {
	switch operator {
	case "+":
		return 2
	case "*":
		return 1
	}
	return -1 // shouldn't get here
}

func infixToPostfix2(expression TokenStream) (result TokenStream) {
	result = make(TokenStream, 0)
	s := stack.New()

	for _, t := range expression {
		switch t.kind {
		case OPERAND:
			result = append(result, t)
		case OPERATOR:
			for s.Len() > 0 && precedence(t.sValue) <= precedence(s.Peek().(Token).sValue) {
				result = append(result, s.Pop().(Token))
			}
			s.Push(t)
		case LPAREN:
			s.Push(t)
		case RPAREN:
			for s.Len() > 0 && s.Peek().(Token).kind != LPAREN {
				result = append(result, s.Pop().(Token))
			}
			_ = s.Pop().(Token) // better be an lparen
		}
	}

	for s.Len() > 0 {
		result = append(result, s.Pop().(Token))
	}

	return
}

// Precondition: expression is in postfix with all parens removed
func eval(expression TokenStream) (value int) {
	s := stack.New()

	for _, token := range expression {
		if token.kind == OPERAND {
			s.Push(token)
		} else {
			op1 := s.Pop().(Token)
			op2 := s.Pop().(Token)
			var result int
			switch token.sValue {
			case "+":
				result = op1.iValue + op2.iValue
			case "*":
				result = op1.iValue * op2.iValue
			}
			s.Push(Token{kind: OPERAND, iValue: result})
		}
	}

	t := s.Pop().(Token)
	value = t.iValue

	return
}

func part1(expressions []TokenStream) {
	sum := 0

	for _, expression := range expressions {
		postfix := infixToPostfix1(expression)
		sum += eval(postfix)
	}

	fmt.Println("Part 1 =", sum)
}

func part2(expressions []TokenStream) {
	sum := 0

	for _, expression := range expressions {
		postfix := infixToPostfix2(expression)
		sum += eval(postfix)
	}

	fmt.Println("Part 2 =", sum)
}

func main() {
	expressions := readData()

	part1(expressions)
	part2(expressions)
}

// Local Variables:
// compile-command: "go build"
// End:
