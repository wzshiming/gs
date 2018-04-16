package gs

import (
	"fmt"
)

type Token int

const (
	_ Token = iota

	ADD // +
	SUB // -
	MUL // *
	QUO // /
)

var tokenMap = map[Token]string{
	ADD: "+",
	SUB: "-",
	MUL: "*",
	QUO: "/",
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
	}
	return 0
}
