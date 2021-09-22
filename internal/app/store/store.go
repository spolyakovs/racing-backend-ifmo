package store

type Store interface {
	User() UserRepository
	Team() TeamRepository
	Driver() DriverRepository
	Race() RaceRepository
}
