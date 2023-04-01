package main

import (
	"fmt"
	"reflect"
)

// const (
// 	default_capacity = 8
// 	LIMIT            = "LIMIT"
// 	FILTER           = "FILTER"
// )

// type Stream[T any] interface {
// }

// func New[T any]() stream[T] {
// 	return func(x T) optional.Optional[T] {
// 		return optional.Of(x)
// 	}
// }

// type operator[T any] struct {
// 	apply func(x T) optional.Optional[T]
// 	cost  float64
// 	name  string
// }

// // func limit[T any](n int) operator[T]

// type stream[T any] func(x T) optional.Optional[T]

// func (s stream[T]) Filter(f func(T) bool) stream[T] {
// 	return func(x T) optional.Optional[T] {
// 		y := s(x)
// 		if y.Empty() {
// 			return y
// 		} else if f(y.Value()) {
// 			return y
// 		}
// 		return optional.Empty[T]()
// 	}
// }

// func (s stream[T]) Limit(n int) stream[T] {
// 	counter := 0
// 	return func(x T) optional.Optional[T] {
// 		if counter > n {
// 			return optional.Empty[T]()
// 		} else if y := s(x); !y.Empty() {
// 			counter++
// 			return y
// 		}
// 		return optional.Empty[T]()
// 	}
// }

// func (s stream[T]) Skip(n int) stream[T] {
// 	counter := 0
// 	return func(x T) optional.Optional[T] {
// 		if counter < n {
// 			counter++
// 			return optional.Empty[T]()
// 		}

// 		return s(x)
// 	}
// }

// func (s stream[T]) ForEach(f func(x T), data []T) {
// 	for _, e := range data {
// 		y := s(e)
// 		if !y.Empty() {
// 			f(e)
// 		}
// 	}
// }

// func (s stream[T]) Reduce(f func(x, y T), data []T) {

// }

// func (s stream[T]) Count(data []T) int {
// 	counter := 0
// 	for _, e := range data {
// 		y := s(e)
// 		if !y.Empty() {
// 			counter++
// 		}
// 	}
// 	return counter
// }

// func t(s []int) {
// 	s[0] = 22
// }

type A struct {
	a string
	b []int
}

func main() {

	// data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// s := New[int]()
	// fmt.Println(s.Filter(func(i int) bool { return i%2 == 0 }).
	// 	Skip(1).
	// 	Limit(2).
	// 	Count(data))

	x := A{
		a: "",
		b: []int{},
	}
	v := reflect.ValueOf(x)

	values := make([]interface{}, v.NumField())

	for i := 0; i < v.NumField(); i++ {
		values[i] = v.Field(i).Interface()
	}

	fmt.Println(values)
}
