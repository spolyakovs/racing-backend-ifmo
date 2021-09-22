package sqlstore

import (
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

type DriverRepository struct {
	store *Store
}

func (driverRepository *DriverRepository) Create(driver *model.Driver) error {
	createDriverQuery := "INSERT INTO drivers (first_name, last_name, birth_date) VALUES ($1, $2, $3) RETURNING id;"
	if err := driverRepository.store.db.QueryRow(
		createDriverQuery,
		driver.FirstName, driver.LastName, driver.BirthDate,
	).Scan(&driver.ID); err != nil {
		return err
	}

	return nil
}

func (driverRepository *DriverRepository) Find(id int) (*model.Driver, error) {
	driver := &model.Driver{}

	findDriverByIDQuery := "SELECT id, first_name, last_name, birth_date FROM drivers WHERE id = $1;"
	if err := driverRepository.store.db.QueryRow(
		findDriverByIDQuery,
		id,
	).Scan(
		&driver.ID,
		&driver.FirstName,
		&driver.LastName,
		&driver.BirthDate,
	); err != nil {
		return nil, err
	}

	return driver, nil
}

func (driverRepository *DriverRepository) Update(driver *model.Driver) error {
	updateDriverQuery := "UPDATE drivers " +
		"SET first_name = $2, " +
		"last_name = $3, " +
		"birth_date = $4 " +
		"WHERE id = $1;"

	countResult, countResultErr := driverRepository.store.db.Exec(
		updateDriverQuery,
		driver.ID,
		driver.FirstName,
		driver.LastName,
		driver.BirthDate,
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

func (driverRepository *DriverRepository) Delete(id int) error {
	deleteDriverQuery := "DELETE FROM drivers WHERE id = $1;"

	countResult, countResultErr := driverRepository.store.db.Exec(
		deleteDriverQuery,
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
