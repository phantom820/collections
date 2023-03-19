package main

import (
	"encoding/json"
	"fmt"

	"github.com/phantom820/collections/sets/hashset"
)

func main() {

	set := hashset.Of("I am the best at what I do and what I do is awesome.", "223")
	// result := function.Map[string](&set,
	// 	func(s string) *vector.Vector[string] {
	// 		vec := vector.Of(strings.Split(s, " ")...)
	// 		return &vec
	// 	})

	// var iterable iterable.Iterable[string] = set.ImmutableCopy()

	// fmt.Println(sets.IsSet(iterable))
	// switch t := iterable.(type) {
	// default:
	// 	fmt.Println(reflect.TypeOf(t))
	// 	fmt.Println(t)

	// }

	b, err := json.Marshal(set)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(b))
}
