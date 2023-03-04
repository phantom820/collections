package errors

import (
	"bytes"
	"errors"
	"text/template"
)

var (
	IndexOutOfBoundsTemplate, _      = template.New("IndexOutOfBounds").Parse("ErrorIndexOutOfBounds: Index {{.index}} out of bounds for length {{.length}}.")
	UnsupportedOperationTemplate, _  = template.New("UnsupportedOperation").Parse("ErrorUnsupportedOperation: Unsupported operation {{.operation}} on [{{.type}}].")
	IndexBoundsOutOfRangeTemplate, _ = template.New("IndexBoundsOutOfRange").Parse("ErrorIndexBoundsOutOfRange: Index bounds [{{.start}}:{{.end}}] out of range.")
	NoSuchElementTemplate, _         = template.New("NoSuchElement").Parse("NoSuchElement: No such element to access.")
)

func IndexOutOfBounds(index int, length int) error {
	var buffer bytes.Buffer
	IndexOutOfBoundsTemplate.Execute(&buffer, map[string]int{"index": index, "length": length})
	return errors.New(buffer.String())
}

func IndexBoundsOutOfRange(start int, end int) error {
	var buffer bytes.Buffer
	IndexBoundsOutOfRangeTemplate.Execute(&buffer, map[string]int{"start": start, "end": end})
	return errors.New(buffer.String())
}

func UnsupportedOperation(operation string, _type string) error {
	var buffer bytes.Buffer
	UnsupportedOperationTemplate.Execute(&buffer, map[string]string{"operation": operation, "type": _type})
	return errors.New(buffer.String())
}

func NoSuchElement() error {
	var buffer bytes.Buffer
	NoSuchElementTemplate.Execute(&buffer, nil)
	return errors.New(buffer.String())
}
