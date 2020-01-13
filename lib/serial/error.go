package serial

import "errors"

var (
	ErrNotStruct = errors.New("not struct")
	ErrConflictPath = errors.New("conflict path")
)

