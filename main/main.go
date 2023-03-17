package main

import (
	"fmt"
	"reflect"

	"github.com/phantom820/collections/iterable"
	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/sets/hashset"
)

func main() {

	set := hashset.Of("I am the best at what I do and what I do is awesome.")
	// result := function.Map[string](&set,
	// 	func(s string) *vector.Vector[string] {
	// 		vec := vector.Of(strings.Split(s, " ")...)
	// 		return &vec
	// 	})

	var iterable iterable.Iterable[string] = set.ImmutableCopy()

	fmt.Println(sets.IsSet(iterable))
	switch t := iterable.(type) {
	default:
		fmt.Println(reflect.TypeOf(t))
		fmt.Println(t)

	}

}
