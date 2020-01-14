package artist

import "errors"

var (
	ErrNotStruct      = errors.New("not struct")
	ErrConflictPath   = errors.New("conflict path")
	ErrNotMultipleOf2 = errors.New("not multiple of 2")
)
