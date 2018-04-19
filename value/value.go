package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

var undefined = fmt.Errorf("Undefined operation")

type Value interface {
	fmt.Stringer
	Binary(t token.Token, y Value) (Value, error)
	PreUnary(t token.Token) (Value, error)
	SufUnary(t token.Token) (Value, error)
}
