package sqlstore

import (
	"database/sql"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

type RaceResultRepository struct {
	store *Store
}

func (raceResultRepository *RaceResultRepository) Create(raceResult *model.RaceResult) error {
	createRaceResultQuery := "INSERT INTO race_results (place, points, race_id, driver_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	if err := raceResultRepository.store.db.QueryRow(
		createRaceResultQuery,
		raceResult.Place,
		raceResult.Points,
		raceResult.Race.ID,
		raceResult.Driver.ID,
	).Scan(&raceResult.ID); err != nil {
		return err
	}

	return nil
}

func (raceResultRepository *RaceResultRepository) Find(id int) (*model.RaceResult, error) {
	findRaceResultByDateQuery := "SELECT id, place, points, race_id, driver_id FROM race_results WHERE id = $1;"

	rows, rowsErr := raceResultRepository.store.db.Query(
		findRaceResultByDateQuery,
		id,
	)

	if rowsErr != nil {
		if rowsErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, rowsErr
	}
	defer rows.Close()

	rows.Next()
	raceResult, raceResultErr := raceResultRepository.scanFromRows(rows)
	if raceResultErr != nil {
		return nil, raceResultErr
	}

	return raceResult, nil
}

func (raceResultRepository *RaceResultRepository) FindByRaceID(race_id int) ([]*model.RaceResult, error) {
	findraceResultByRaceIDQuery := "SELECT id, place, points, race_id, driver_id FROM race_results WHERE race_id = $1;"

	rows, rowsErr := raceResultRepository.store.db.Query(
		findraceResultByRaceIDQuery,
		race_id,
	)
	if rowsErr != nil {
		if rowsErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, rowsErr
	}
	defer rows.Close()

	var results []*model.RaceResult

	for rows.Next() {
		raceResult, err := raceResultRepository.scanFromRows(rows)
		if err != nil {
			return results, err
		}

		results = append(results, raceResult)
	}

	if err := rows.Err(); err != nil {
		return results, err
	}

	return results, nil
}

func (raceResultRepository *RaceResultRepository) FindByDriverID(driver_id int) ([]*model.RaceResult, error) {
	findraceResultByDriverIDQuery := "SELECT id, place, points, race_id, driver_id FROM race_results WHERE driver_id = $1;"

	rows, rowsErr := raceResultRepository.store.db.Query(
		findraceResultByDriverIDQuery,
		driver_id,
	)
	if rowsErr != nil {
		if rowsErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, rowsErr
	}
	defer rows.Close()

	var results []*model.RaceResult

	for rows.Next() {
		raceResult, err := raceResultRepository.scanFromRows(rows)
		if err != nil {
			return results, err
		}

		results = append(results, raceResult)
	}

	if err := rows.Err(); err != nil {
		return results, err
	}

	return results, nil
}

func (raceResultRepository *RaceResultRepository) Update(raceResult *model.RaceResult) error {
	updateRaceResultQuery := "UPDATE race_results " +
		"SET place = $2, " +
		"points = $3, " +
		"race_id = $4 " +
		"driver_id = $5, " +
		"WHERE id = $1;"

	countResult, countResultErr := raceResultRepository.store.db.Exec(
		updateRaceResultQuery,
		raceResult.ID,
		raceResult.Place,
		raceResult.Points,
		raceResult.Race.ID,
		raceResult.Driver.ID,
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

func (raceResultRepository *RaceResultRepository) Delete(id int) error {
	deleteRaceResultQuery := "DELETE FROM race_results WHERE id = $1;"

	countResult, countResultErr := raceResultRepository.store.db.Exec(
		deleteRaceResultQuery,
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

func (raceResultRepository *RaceResultRepository) scanFromRows(rows *sql.Rows) (*model.RaceResult, error) {
	var (
		raceResult *model.RaceResult
		raceID     int
		driverID   int
	)

	if err := rows.Scan(
		&raceResult.ID,
		&raceResult.Place,
		&raceResult.Points,
		&raceID,
		&driverID,
	); err != nil {
		return nil, err
	}

	tmpRace, raceErr := raceResultRepository.store.Race().Find(raceID)
	if raceErr != nil {
		return nil, raceErr
	}

	tmpDriver, driverErr := raceResultRepository.store.Driver().Find(driverID)
	if driverErr != nil {
		return nil, driverErr
	}

	raceResult.Race = tmpRace
	raceResult.Driver = tmpDriver

	return raceResult, nil
}
