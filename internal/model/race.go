package model

import "time"

type Race struct {
	ID       int
	Name     string
	Location string
	Date     time.Time
}
