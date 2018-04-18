package gs

import (
	"fmt"
)

type Token uint

const (
	INVALID Token = iota

	NUMBER // 123 or 123.4
	IDENT  // abc or a123

	operatorBeg
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	PERIOD // .
	COMMA  // ,

	LPAREN // (
	LBRACK // [
	LBRACE // {

	RPAREN // )
	RBRACK // ]
	RBRACE // }
	operatorEnd

	keyworkBeg
	IF   // if
	ELSE // else
	keyworkEnd
)

var tokenMap = map[Token]string{
	NUMBER: "number",
	IDENT:  "ident",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	QUO:    "/",
	PERIOD: ".",
	COMMA:  ",",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",

	RPAREN: ")",
	RBRACK: "]",
	RBRACE: "}",

	IF:   "if",
	ELSE: "else",
}

func (op Token) String() string {
	v, ok := tokenMap[op]
	if ok {
		return v
	}
	return fmt.Sprintf("Token(%d)", op)
}

var prec = map[Token]int{}

func init() {
	ADD.SetPrecedence(2)
	SUB.SetPrecedence(2)
	MUL.SetPrecedence(3)
	QUO.SetPrecedence(3)
	PERIOD.SetPrecedence(4)
	COMMA.SetPrecedence(5)
}

func (op Token) SetPrecedence(pre int) {
	prec[op] = pre
}

func (op Token) Precedence() int {
	return prec[op]
}

var ks = map[string]Token{}
var os = map[string]Token{}

func init() {
	for i := keyworkBeg; i != keyworkEnd; i++ {
		ks[tokenMap[i]] = i
	}
	for i := operatorBeg; i != operatorEnd; i++ {
		os[tokenMap[i]] = i
	}
}

func LookupKeywork(s string) Token {
	return ks[s]
}

func LookupOperator(s string) Token {
	return os[s]
}

func (t Token) IsKeywork() bool {
	return keyworkBeg < t && t < keyworkEnd
}

func (t Token) IsOperator() bool {
	return operatorBeg < t && t < operatorEnd
}
