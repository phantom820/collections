package main

import (
	"collections/set/hashset"
	"collections/wrapper"
	"fmt"
)

func main() {
	// // Haha my own collections.
	s := hashset.NewHashSet[wrapper.Integer]()
	s.Add(1)
	s.Add(1)

	fmt.Println(s)
}
