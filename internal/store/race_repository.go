package store

import (
	"time"

	"github.com/spolyakovs/racing-backend-ifmo/internal/model"
)

type RaceRepository struct {
	store *Store
}

func (raceRepository *RaceRepository) Create(race *model.Race) (*model.Race, error) {
	create_race_query := "INSERT INTO races (name, location, date) VALUES ($1, $2, $3) RETURNING id"
	if err := raceRepository.store.db.QueryRow(
		create_race_query,
		race.Name, race.Location, race.Date,
	).Scan(&race.ID); err != nil {
		return nil, err
	}

	return race, nil
}

func (raceRepository *RaceRepository) FindByDate(date time.Time) (*model.Race, error) {
	race := &model.Race{}
	if err := raceRepository.store.db.QueryRow(
		"SELECT id, name, location, date FROM races WHERE date = $1",
		date,
	).Scan(
		&race.ID,
		&race.Name,
		&race.Location,
		&race.Date,
	); err != nil {
		return nil, err
	}

	return race, nil
}

func (raceRepository *RaceRepository) Update(race *model.Race) (*model.Race, error) {
	return nil, nil
}
