package store

import (
	"database/sql"
)

type Store struct {
	db                           *sql.DB
	userRepository               *UserRepository
	teamRepository               *TeamRepository
	driverRepository             *DriverRepository
	raceRepository               *RaceRepository
	teamDriverContractRepository *TeamDriverContractRepository
	raceResultRepository         *RaceResultRepository
}

// TODO: unite into one functions Create(except User), Find..., Update and Delete, probably in "store.go", args: (query string, queryVars ...interface{})
func New(db *sql.DB) (*Store, error) {
	newStore := &Store{
		db: db,
	}

	if err := newStore.createTables(); err != nil {
		return nil, err
	}

	return newStore, nil
}

func (st *Store) User() *UserRepository {
	if st.userRepository != nil {
		return st.userRepository
	}

	st.userRepository = &UserRepository{
		DefaultRepository{
			store: st,
		},
	}

	return st.userRepository
}

func (st *Store) Team() *TeamRepository {
	if st.teamRepository != nil {
		return st.teamRepository
	}

	st.teamRepository = &TeamRepository{
		store: st,
	}

	return st.teamRepository
}

func (st *Store) Driver() *DriverRepository {
	if st.driverRepository != nil {
		return st.driverRepository
	}

	st.driverRepository = &DriverRepository{
		store: st,
	}

	return st.driverRepository
}

func (st *Store) Race() *RaceRepository {
	if st.raceRepository != nil {
		return st.raceRepository
	}

	st.raceRepository = &RaceRepository{
		store: st,
	}

	return st.raceRepository
}

func (st *Store) TeamDriverContract() *TeamDriverContractRepository {
	if st.teamDriverContractRepository != nil {
		return st.teamDriverContractRepository
	}

	st.teamDriverContractRepository = &TeamDriverContractRepository{
		store: st,
	}

	return st.teamDriverContractRepository
}

func (st *Store) RaceResult() *RaceResultRepository {
	if st.raceResultRepository != nil {
		return st.raceResultRepository
	}

	st.raceResultRepository = &RaceResultRepository{
		store: st,
	}

	return st.raceResultRepository
}
