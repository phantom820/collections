package main

import (
	"collections/list"
	"collections/types"
	"fmt"
)

func main() {
	// // Haha my own collections.
	l := list.NewList[types.Integer]()

	fmt.Println(l)
	r := l.Reverse()
	fmt.Println(r)
	r.RemoveFront()
	fmt.Println(r)
	fmt.Println(l)

}
