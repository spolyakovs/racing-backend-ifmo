package model

type RaceResult struct {
	ID     int
	Place  int
	Points int
	Race   *Race
	Driver *Driver
}
