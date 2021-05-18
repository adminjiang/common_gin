package lib

import "errors"

type Errors []error

var (
	// ErrRecordNotFound record not found error, happens when only haven't find any matched data when looking up with a struct, finding a slice won't return this error
	ErrRecordNotFound = errors.New("record not found")
)