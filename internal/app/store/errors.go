package store

import "errors"

var (
	ErrWrongTable     = errors.New("Wrong table name")
	ErrWrongModel     = errors.New("Given wrong model")
	ErrRecordNotFound = errors.New("Record not found")
)
