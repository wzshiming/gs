package value

import (
	"fmt"

	"github.com/wzshiming/gs/token"
)

type ValueVar struct {
	Scope *Scope
	Name  string
}

func (v *ValueVar) String() string {
	if v == nil {
		return "<ValueVar.nil>"
	}
	val, ok := v.Scope.Get(v.Name)
	if !ok || val == nil {
		return "<ValueVar.nil>"
	}
	return val.String()
}

func (v *ValueVar) Point() (Value, error) {
	val, ok := v.Scope.Get(v.Name)
	if !ok {
		return v, fmt.Errorf("Variable is empty '%v'", v.Name)
	}
	return val, nil
}

func (v *ValueVar) Binary(t token.Token, y Value) (Value, error) {

	switch t {
	case token.ASSIGN:
		v.Scope.Set(v.Name, y)
		return v, nil
	case token.DEFINE:
		v.Scope.SetLocal(v.Name, y)
		return v, nil
	}

	val, err := v.Point()
	if err != nil {
		return v, err
	}

	return val.Binary(t, y)
}

func (v *ValueVar) UnaryPre(t token.Token) (Value, error) {
	return v, undefined
}

func (v *ValueVar) UnarySuf(t token.Token) (Value, error) {
	return v, undefined
}
