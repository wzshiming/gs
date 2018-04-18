package token

import (
	"fmt"
)

type Token uint

const (
	INVALID Token = iota

	STRING // "123" or '123'
	NUMBER // 123 or 123.4
	IDENT  // abc or a123

	operatorBeg
	ADD    // +
	SUB    // -
	MUL    // *
	QUO    // /
	POW    // **
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
	STRING: "string",
	NUMBER: "number",
	IDENT:  "ident",

	ADD:    "+",
	SUB:    "-",
	MUL:    "*",
	QUO:    "/",
	POW:    "**",
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
	ADD.setPrecedence(2)
	SUB.setPrecedence(2)
	MUL.setPrecedence(3)
	QUO.setPrecedence(3)
	POW.setPrecedence(4)
	PERIOD.setPrecedence(5)
	COMMA.setPrecedence(6)
}

func (op Token) setPrecedence(pre int) {
	prec[op] = pre
}

func (op Token) Precedence() int {
	return prec[op]
}

func (t Token) IsKeywork() bool {
	return keyworkBeg < t && t < keyworkEnd
}

func (t Token) IsOperator() bool {
	return operatorBeg < t && t < operatorEnd
}

var LookupKeywork = newLooker()
var LookupOperator = newLooker()

func init() {
	for i := keyworkBeg + 1; i != keyworkEnd; i++ {
		LookupKeywork.Add([]rune(tokenMap[i]), i)
	}
	for i := operatorBeg + 1; i != operatorEnd; i++ {
		LookupOperator.Add([]rune(tokenMap[i]), i)
	}
}
