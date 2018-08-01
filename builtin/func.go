package builtin

import (
	"fmt"

	"github.com/wzshiming/gs/value"
)

var Func = map[string]interface{}{
	"println":  fmt.Println,
	"print":    fmt.Print,
	"sprintln": fmt.Sprintln,
	"sprint":   fmt.Sprint,
	"string":   funcString,
	"number":   funcNumber,
}

func funcString(val value.Value) value.Value {
	switch val.(type) {
	case *value.Var:
		return funcString(val.Point())
	case value.String:
		return val
	default:
		return value.String(val.String())
	}
}

func funcNumber(val value.Value) value.Value {
	switch val.(type) {
	case *value.Var:
		return funcNumber(val.Point())
	case value.Number:
		return val
	default:
		return value.ParseNumber(val.String())
	}
}
