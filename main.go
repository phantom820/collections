package main

import (
	"collections/set"
	"collections/wrapper"
	"fmt"
)

func main() {
	// // Haha my own collections.
	s := set.NewHashSet[wrapper.Integer]()
	s.Add(1)
	s.Add(1)

	fmt.Println(s)
}
