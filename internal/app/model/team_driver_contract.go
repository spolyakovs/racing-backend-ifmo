package model

import (
	"time"
)

type TeamDriverContract struct {
	ID       int
	FromDate time.Time
	ToDate   time.Time
	Team     *Team
	Driver   *Driver
}
