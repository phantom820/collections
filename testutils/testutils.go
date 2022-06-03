// Package testutils provides some utility functions for writing unit tests for collections such as checking equality of slices using corresponding elements.
package testutils

import "github.com/phantom820/collections/types"

func EqualSlices[T types.Equitable[T]](a []T, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}
	return true
}
