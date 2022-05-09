package main

import (
	"collections/list"
	"collections/wrapper"
	"fmt"
)

func main() {
	// // Haha my own collections.
	l := list.NewList[wrapper.Integer]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	fmt.Println(l)
	r := l.Reverse()
	fmt.Println(r)
	r.RemoveFront()
	fmt.Println(r)
	fmt.Println(l)

}
