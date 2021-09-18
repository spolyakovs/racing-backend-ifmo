package model

import "time"

type Driver struct {
	ID        int
	FirstName string
	LastName  string
	BirthDate time.Time
}
