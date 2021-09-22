package model

import "errors"

var (
	ErrGetFields      = errors.New("Cannot get fields from default model")
	ErrWrongModel     = errors.New("Given wrong model")
	ErrRecordNotFound = errors.New("Record not found")
)
