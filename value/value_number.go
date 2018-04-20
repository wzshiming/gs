package value

import (
	"math"
	"strconv"
	"strings"
)

const (
	maxInt   = math.MaxInt32
	minInt   = math.MinInt32
	maxFloat = math.MaxFloat32
	minFloat = math.SmallestNonzeroFloat32
)

type ValueNumber interface {
	Value
	Int() valueNumberInt
	Float() valueNumberFloat
	BigInt() valueNumberBigInt
	BigFloat() valueNumberBigFloat
}

func ParseValueNumber(s string) ValueNumber {
	if strings.Index(s, ".") != -1 {
		val, _ := strconv.ParseFloat(s, 0)
		return newValueNumberFloat(val)
	} else {
		val, _ := strconv.ParseInt(s, 0, 0)
		return newValueNumberInt(val)
	}
}
