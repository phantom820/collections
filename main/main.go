package main

import (
	"github.com/phantom820/collections/sets"
	"github.com/phantom820/collections/sets/hashset"
	"github.com/phantom820/collections/types"
)

func main() {

	s1 := hashset.New[types.Int]()
	// s2 :=
	sets.Union[types.Int](s1, s1)
}
