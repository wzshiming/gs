package builtin

import (
	"fmt"
)

var Func = map[string]interface{}{
	"println":  fmt.Println,
	"print":    fmt.Print,
	"sprintln": fmt.Sprintln,
	"sprint":   fmt.Sprint,
}
