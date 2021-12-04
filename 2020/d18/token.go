package main

import "fmt"

const (
	LPAREN   = iota
	RPAREN   = iota
	OPERAND  = iota
	OPERATOR = iota
)

type Token struct {
	kind   int
	iValue int
	sValue string
}

func (t Token) String() (result string) {
	var kind string
	switch t.kind {
	case LPAREN:
		kind = "LP"
	case RPAREN:
		kind = "RP"
	case OPERAND:
		kind = "VL"
	case OPERATOR:
		kind = "OP"
	}

	if kind == "LP" || kind == "RP" {
		result = fmt.Sprintf("<%s>", kind)
	} else if kind == "OP" {
		result = fmt.Sprintf("<%s:%s>", kind, t.sValue)
	} else {
		result = fmt.Sprintf("<%s: %d>", kind, t.iValue)
	}

	return result
}

type TokenStream []Token
