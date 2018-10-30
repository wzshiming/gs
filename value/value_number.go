package value

import (
	"strconv"
	"strings"
)

type Number interface {
	Value
	Int() numberInt
	Float() numberFloat
}

func ParseNumber(s string) Number {
	if strings.Index(s, ".") != -1 {
		val, _ := strconv.ParseFloat(s, 0)
		return numberFloat(val)
	}
	val, _ := strconv.ParseInt(s, 0, 0)
	return numberInt(val)
}
