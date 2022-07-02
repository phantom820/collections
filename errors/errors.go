// package errors provides custom errors to be used
package errors

import (
	"bytes"
	"html/template"
)

// error codes.
const (
	IndexOutOfBounds = 1
	NoSuchElement    = 2
	NoNextElement    = 3
	MapKeyRange      = 4
)

// error templates.
var (
	IndexOutOfBoundsTemplate, _ = template.New("IndexOutOfBounds").Parse("ErrIndexOutOfBounds: Index: {{.index}}, Size: {{.size}}")
	NoSuchElementTemplate, _    = template.New("NoSuchElement").Parse("ErrNoSuchElement , Size: {{.size}}")
	NoNextElementTemplate, _    = template.New("NoNextElement").Parse("ErrNoNextElement")
	MapKeyRangeTemplate, _      = template.New("MapKeyRange").Parse("ErrMapKeyRange  lower key: {{.fromKey}} , upper Key: {{.toKey}}, lower key greater than upper key.")
)

// Error a custom error type to be used.
type Error struct {
	code int
	msg  string
	Err  error
}

// Code returns the error code for the error.
func (err Error) Code() int {
	return err.code
}

// Error returns the error message.
func (err *Error) Error() string {
	return err.msg
}

func ErrIndexOutOfBounds(index int, size int) Error {
	var buffer bytes.Buffer
	IndexOutOfBoundsTemplate.Execute(&buffer, map[string]int{"index": index, "size": size})
	return Error{code: IndexOutOfBounds, msg: buffer.String()}
}

func ErrNoSuchElement(size int) Error {
	var buffer bytes.Buffer
	NoSuchElementTemplate.Execute(&buffer, map[string]int{"size": size})
	return Error{code: NoSuchElement, msg: buffer.String()}
}

func ErrNoNextElement() Error {
	var buffer bytes.Buffer
	NoNextElementTemplate.Execute(&buffer, map[string]int{})
	return Error{code: NoNextElement, msg: buffer.String()}
}

func ErrMapKeyRange[T any](fromKey T, toKey T) Error {
	var buffer bytes.Buffer
	MapKeyRangeTemplate.Execute(&buffer, map[string]T{"fromKey": fromKey, "toKey": toKey})
	return Error{code: MapKeyRange, msg: buffer.String()}
}
