// errors defines custom error type [Error] to be used for handling errors on collections.
package errors

import (
	"bytes"
	"errors"
	"text/template"
)

const (
	IndexOutOfBoundsCode      = 1 //  Invalid indexing i.e indexing a buffere outside its range.
	UnsupportedOperationCode  = 2 //  An operation that is not supported i.e mutating operation on an immutable data structure.
	IndexBoundsOutOfRangeCode = 3 //  Misconfigured indexing i.e lower index being greater than upper index.
	NoSuchElementCode         = 4 //  An absent element i.e next on iterator without has next guard.
)

var (
	indexOutOfBoundsTemplate, _      = template.New("IndexOutOfBounds").Parse("ErrorIndexOutOfBounds: Index {{.index}} out of bounds for length {{.length}}.")
	unsupportedOperationTemplate, _  = template.New("UnsupportedOperation").Parse("ErrorUnsupportedOperation: Unsupported operation {{.operation}} on [{{.type}}].")
	indexBoundsOutOfRangeTemplate, _ = template.New("IndexBoundsOutOfRange").Parse("ErrorIndexBoundsOutOfRange: Index bounds [{{.start}}:{{.end}}] out of range.")
	noSuchElementTemplate, _         = template.New("NoSuchElement").Parse("NoSuchElement: No such element to access.")
)

// Error custom error type for collections.
type Error struct {
	code  int // The error code.
	error     // The actual underlying error.
}

// New creates an error with the given code and underlying error.
func New(code int, err error) Error {
	return Error{code: code, error: err}
}

func IndexOutOfBounds(index int, length int) Error {
	var buffer bytes.Buffer
	indexOutOfBoundsTemplate.Execute(&buffer, map[string]int{"index": index, "length": length})
	return New(IndexOutOfBoundsCode, errors.New(buffer.String()))
}

func IndexBoundsOutOfRange(start int, end int) Error {
	var buffer bytes.Buffer
	indexBoundsOutOfRangeTemplate.Execute(&buffer, map[string]int{"start": start, "end": end})
	return New(IndexBoundsOutOfRangeCode, errors.New(buffer.String()))
}

func UnsupportedOperation(operation string, _type string) Error {
	var buffer bytes.Buffer
	unsupportedOperationTemplate.Execute(&buffer, map[string]string{"operation": operation, "type": _type})
	return New(UnsupportedOperationCode, errors.New(buffer.String()))
}

func NoSuchElement() Error {
	var buffer bytes.Buffer
	noSuchElementTemplate.Execute(&buffer, nil)
	return New(NoSuchElementCode, errors.New(buffer.String()))
}
