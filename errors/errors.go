package errors

import (
	"bytes"
	"fmt"

	"github.com/wzshiming/gs/position"
)

type Error struct {
	Pos position.Position
	Err error
}

func (e *Error) String() string {
	return fmt.Sprintf("%v %s\n", e.Pos, e.Err.Error())
}

func (e *Error) Error() string {
	return e.String()
}

func NewErrors() *Errors {
	return &Errors{}
}

type Errors []*Error

func (e *Errors) Append(pos position.Position, err error) {
	*e = append(*e, &Error{
		Pos: pos,
		Err: err,
	})
}

func (e *Errors) String() string {
	buf := bytes.NewBuffer(nil)
	for _, v := range *e {
		buf.WriteString(v.String())
	}
	return buf.String()
}

func (e *Errors) Error() string {
	return e.String()
}

func (e *Errors) Len() int {
	return len(*e)
}
