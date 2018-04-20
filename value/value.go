package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

var undefined = fmt.Errorf("Undefined operation")

type Value interface {
	fmt.Stringer
	Binary(t token.Token, y Value) (Value, error)
	UnaryPre(t token.Token) (Value, error)
	UnarySuf(t token.Token) (Value, error)
}
