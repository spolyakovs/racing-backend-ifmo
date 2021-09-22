package store

import (
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
)

type TeamRepository struct {
	store *Store
}

func (teamRepository *TeamRepository) Create(team *model.Team) error {
	createTeamQuery := "INSERT INTO teams (name, engine_manufacturer) VALUES ($1, $2) RETURNING id"
	if err := teamRepository.store.db.QueryRow(
		createTeamQuery,
		team.Name, team.EngineManufacturer,
	).Scan(&team.ID); err != nil {
		return err
	}

	return nil
}

func (teamRepository *TeamRepository) Find(id int) (*model.Team, error) {
	team := &model.Team{}

	findTeamByIDQuery := "SELECT id, name, engine_manufacturer FROM teams WHERE id = $1"
	if err := teamRepository.store.db.QueryRow(
		findTeamByIDQuery,
		id,
	).Scan(
		&team.ID,
		&team.Name,
		&team.EngineManufacturer,
	); err != nil {
		return nil, err
	}

	return team, nil
}

func (teamRepository *TeamRepository) FindByName(name string) (*model.Team, error) {
	team := &model.Team{}

	findTeamByNameQuery := "SELECT id, name, engine_manufacturer FROM teams WHERE name = $1"
	if err := teamRepository.store.db.QueryRow(
		findTeamByNameQuery,
		name,
	).Scan(
		&team.ID,
		&team.Name,
		&team.EngineManufacturer,
	); err != nil {
		return nil, err
	}

	return team, nil
}

func (teamRepository *TeamRepository) Update(team *model.Team) error {
	updateTeamQuery := "UPDATE teams " +
		"SET name = $2, " +
		"engine_manufacturer = $3 " +
		"WHERE id = $1"

	countResult, countResultErr := teamRepository.store.db.Exec(
		updateTeamQuery,
		team.ID,
		team.Name,
		team.EngineManufacturer,
	)

	if countResultErr != nil {
		return countResultErr
	}

	count, countErr := countResult.RowsAffected()

	if countErr != nil {
		return countErr
	}

	if count == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (teamRepository *TeamRepository) Delete(id int) error {
	deleteTeamQuery := "DELETE FROM teams WHERE id = $1"

	countResult, countResultErr := teamRepository.store.db.Exec(
		deleteTeamQuery,
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
		return ErrRecordNotFound
	}

	return nil
}
