package gs

import (
	"fmt"
)

type Token int

const (
	_ Token = iota

	NUMBER // 123 or 123.4
	IDENT  // abc or a123

	ADD // +
	SUB // -
	MUL // *
	QUO // /
	DOT // .
)

var tokenMap = map[Token]string{
	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
	DOT: ".",
}

func (op Token) String() string {
	v, ok := tokenMap[op]
	if ok {
		return v
	}
	return fmt.Sprintf("Token(%d)", op)
}

func (op Token) Precedence() int {
	switch op {
	case ADD, SUB:
		return 2
	case MUL, QUO:
		return 3
	case DOT:
		return 4
	}
	return 0
}
