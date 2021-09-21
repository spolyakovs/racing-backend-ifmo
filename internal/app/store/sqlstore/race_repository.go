package sqlstore

import (
	"time"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

type RaceRepository struct {
	store *Store
}

func (raceRepository *RaceRepository) Create(race *model.Race) error {
	createRaceQuery := "INSERT INTO races (name, location, date) VALUES ($1, $2, $3) RETURNING id"
	if err := raceRepository.store.db.QueryRow(
		createRaceQuery,
		race.Name, race.Location, race.Date,
	).Scan(&race.ID); err != nil {
		return err
	}

	return nil
}

func (raceRepository *RaceRepository) Find(id int) (*model.Race, error) {
	race := &model.Race{}

	findRaceByIDQuery := "SELECT id, name, location, date FROM races WHERE id = $1"
	if err := raceRepository.store.db.QueryRow(
		findRaceByIDQuery,
		id,
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

func (raceRepository *RaceRepository) FindByDate(date time.Time) (*model.Race, error) {
	race := &model.Race{}

	findRaceByDateQuery := "SELECT id, name, location, date FROM races WHERE date = $1"
	if err := raceRepository.store.db.QueryRow(
		findRaceByDateQuery,
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

func (raceRepository *RaceRepository) Update(race *model.Race) error {
	updateRaceQuery := "UPDATE races " +
		"SET name = $2, " +
		"location = $3, " +
		"date = $4 " +
		"WHERE id = $1"

	countResult, countResultErr := raceRepository.store.db.Exec(
		updateRaceQuery,
		race.ID,
		race.Name,
		race.Location,
		race.Date,
	)

	if countResultErr != nil {
		return countResultErr
	}

	count, countErr := countResult.RowsAffected()

	if countErr != nil {
		return countErr
	}

	if count == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}

func (raceRepository *RaceRepository) Delete(id int) error {
	deleteRaceQuery := "DELETE FROM races WHERE id = $1"

	countResult, countResultErr := raceRepository.store.db.Exec(
		deleteRaceQuery,
		id,
	)

	if countResultErr != nil {
		return countResultErr
	}

	count, countErr := countResult.RowsAffected()

	if countErr != nil {
		return countErr
	}

	if count == 0 {
		return store.ErrRecordNotFound
	}

	return nil
}
