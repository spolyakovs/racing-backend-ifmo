package store

import "github.com/spolyakovs/racing-backend-ifmo/internal/model"

type TeamRepository struct {
	store *Store
}

func (teamRepository *TeamRepository) Create(team *model.Team) (*model.Team, error) {
	create_team_query := "INSERT INTO teams (name, engine_manufacturer) VALUES ($1, $2) RETURNING id"
	if err := teamRepository.store.db.QueryRow(
		create_team_query,
		team.Name, team.EngineManufacturer,
	).Scan(&team.ID); err != nil {
		return nil, err
	}

	return team, nil
}

func (teamRepository *TeamRepository) FindByName(name string) (*model.Team, error) {
	team := &model.Team{}
	if err := teamRepository.store.db.QueryRow(
		"SELECT id, name, engine_manufacturer FROM teams WHERE name = $1",
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

func (teamRepository *TeamRepository) Update(team *model.Team) (*model.Team, error) {
	return nil, nil
}
