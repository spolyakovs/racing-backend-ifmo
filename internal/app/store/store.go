package store

type Store interface {
	Users() UserRepository
	Teams() TeamRepository
	Drivers() DriverRepository
	Races() RaceRepository
	TeamDriverContracts() TeamDriverContractRepository
	RaceResults() RaceResultRepository
}
