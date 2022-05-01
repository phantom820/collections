package wrapper

import (
	"collections/interfaces"
	"math"
)

type Integer int

func (i Integer) HashCode() int {
	return int(i)
}

func (i Integer) Equals(other Integer) bool {
	return i == other
}

type String string

func (s String) HashCode() int {
	runes := []rune(s)
	code := 0
	for i, r := range runes {
		v := float64(int(r))
		code = code + int(v*(math.Pow(2.0, float64(len(runes)-i))))
	}
	return code
}

func (s String) Equals(other String) bool {
	return s == other
}

func (s String) Less(other String) bool {
	return s < other
}

func (x Integer) Less(y Integer) bool {
	return x < y
}

type Slice[T interfaces.Equitable[T]] []T

func (s Slice[T]) Collect() []T {
	slice := []T(s)
	return slice
}
