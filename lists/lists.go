package lists

import "errors"

// Errors for  operations that may be invalid on a list.
var (
	ErrEmptyList   = errors.New("cannot remove from an empty list")
	ErrOutOfBounds = errors.New("index out of bounds")
)
