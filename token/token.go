package token

import (
	"fmt"
)

type Token uint

const (
	INVALID Token = iota
	COMMENT
	EOF

	literalBeg
	NIL    // nil
	BOOL   // true or false
	STRING // "123" or '123' or `123`
	NUMBER // 123 or 123.4
	IDENT  // abc or a123
	literalEnd

	operatorBeg
	ADD // +
	SUB // -
	MUL // *
	QUO // /
	POW // **
	REM // %

	AND     // &
	OR      // |
	XOR     // ^
	SHL     // <<
	SHR     // >>
	AND_NOT // &^

	ADD_ASSIGN // +=
	SUB_ASSIGN // -=
	MUL_ASSIGN // *=
	QUO_ASSIGN // /=
	POW_ASSIGN // **=
	REM_ASSIGN // %=

	AND_ASSIGN     // &=
	OR_ASSIGN      // |=
	XOR_ASSIGN     // ^=
	SHL_ASSIGN     // <<=
	SHR_ASSIGN     // >>=
	AND_NOT_ASSIGN // &^=

	LAND  // &&
	LOR   // ||
	ARROW // <-
	INC   // ++
	DEC   // --

	EQL    // ==
	LSS    // <
	GTR    // >
	ASSIGN // =
	NOT    // !

	NEQ      // !=
	LEQ      // <=
	GEQ      // >=
	DEFINE   // :=
	ELLIPSIS // ...

	LPAREN // (
	LBRACK // [
	LBRACE // {
	COMMA  // ,
	PERIOD // .

	RPAREN    // )
	RBRACK    // ]
	RBRACE    // }
	SEMICOLON // ;
	COLON     // :
	operatorEnd

	keyworkBeg
	MAP    // map
	IF     // if
	ELSE   // else
	FOR    // for
	FUNC   // func
	RETURN // return
	keyworkEnd
)

var tokenMap = map[Token]string{
	INVALID: "invalid",
	EOF:     "eof",

	NIL:    "nil",
	BOOL:   "bool",
	STRING: "string",
	NUMBER: "number",
	IDENT:  "ident",

	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	POW: "**",
	REM: "%",

	AND:     "&",
	OR:      "|",
	XOR:     "^",
	SHL:     "<<",
	SHR:     ">>",
	AND_NOT: "&^",

	ADD_ASSIGN: "+=",
	SUB_ASSIGN: "-=",
	MUL_ASSIGN: "*=",
	QUO_ASSIGN: "/=",
	POW_ASSIGN: "**=",
	REM_ASSIGN: "%=",

	AND_ASSIGN:     "&=",
	OR_ASSIGN:      "|=",
	XOR_ASSIGN:     "^=",
	SHL_ASSIGN:     "<<=",
	SHR_ASSIGN:     ">>=",
	AND_NOT_ASSIGN: "&^=",

	LAND:  "&&",
	LOR:   "||",
	ARROW: "<-",
	INC:   "++",
	DEC:   "--",

	EQL:    "==",
	LSS:    "<",
	GTR:    ">",
	ASSIGN: "=",
	NOT:    "!",

	NEQ:      "!=",
	LEQ:      "<=",
	GEQ:      ">=",
	DEFINE:   ":=",
	ELLIPSIS: "...",

	LPAREN: "(",
	LBRACK: "[",
	LBRACE: "{",
	COMMA:  ",",
	PERIOD: ".",

	RPAREN:    ")",
	RBRACK:    "]",
	RBRACE:    "}",
	SEMICOLON: ";",
	COLON:     ":",

	MAP:    "map",
	IF:     "if",
	ELSE:   "else",
	FOR:    "for",
	FUNC:   "func",
	RETURN: "return",
}

func (op Token) String() string {
	v, ok := tokenMap[op]
	if ok {
		return v
	}
	return fmt.Sprintf("Token(%d)", op)
}

var prec = map[Token]int{}

func (op Token) setPrecedence(pre int) {
	prec[op] = pre
}

func (op Token) Precedence() int {
	return prec[op]
}

func (t Token) IsLiteral() bool {
	return literalBeg < t && t < literalEnd
}

func (t Token) IsKeywork() bool {
	return keyworkBeg < t && t < keyworkEnd
}

func (t Token) IsOperator() bool {
	return operatorBeg < t && t < operatorEnd
}
