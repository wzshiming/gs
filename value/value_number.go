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

type Number interface {
	Value
	Int() numberInt
	Float() numberFloat
	BigInt() numberBigInt
	BigFloat() numberBigFloat
}

func ParseValueNumber(s string) Number {
	if strings.Index(s, ".") != -1 {
		val, _ := strconv.ParseFloat(s, 0)
		return newNumberFloat(val)
	}
	val, _ := strconv.ParseInt(s, 0, 0)
	return newNumberInt(val)
}
