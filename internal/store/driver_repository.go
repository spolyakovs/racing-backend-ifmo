package store

import "github.com/spolyakovs/racing-backend-ifmo/internal/model"

type DriverRepository struct {
	store *Store
}

func (driverRepository *DriverRepository) Create(driver *model.Driver) (*model.Driver, error) {
	create_driver_query := "INSERT INTO drivers (first_name, last_name, birth_date) VALUES ($1, $2, $3) RETURNING id"
	if err := driverRepository.store.db.QueryRow(
		create_driver_query,
		driver.FirstName, driver.LastName, driver.BirthDate,
	).Scan(&driver.ID); err != nil {
		return nil, err
	}

	return driver, nil
}

func (driverRepository *DriverRepository) FindByID(id int) (*model.Driver, error) {
	driver := &model.Driver{}
	if err := driverRepository.store.db.QueryRow(
		"SELECT id, first_name, last_name, birth_date FROM drivers WHERE id = $1",
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

func (driverRepository *DriverRepository) Update(driver *model.Driver) (*model.Driver, error) {
	return nil, nil
}
