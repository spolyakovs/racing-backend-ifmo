package sqlstore

import (
	"database/sql"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

type TeamDriverContractRepository struct {
	store *Store
}

func (teamDriverContractRepository *TeamDriverContractRepository) Create(teamDriverContract *model.TeamDriverContract) error {
	createTeamDriverContractQuery := "INSERT INTO team_driver_contracts (from_date, to_date, team_id, driver_id) VALUES ($1, $2, $3, $4) RETURNING id;"
	if err := teamDriverContractRepository.store.db.QueryRow(
		createTeamDriverContractQuery,
		teamDriverContract.FromDate,
		teamDriverContract.ToDate,
		teamDriverContract.Team.ID,
		teamDriverContract.Driver.ID,
	).Scan(&teamDriverContract.ID); err != nil {
		return err
	}

	return nil
}

func (teamDriverContractRepository *TeamDriverContractRepository) Find(id int) (*model.TeamDriverContract, error) {
	findTeamDriverContractByIDQuery := "SELECT id, from_date, to_date, team_id, driver_id FROM team_driver_contracts WHERE id = $1;"
	rows, rowsErr := teamDriverContractRepository.store.db.Query(
		findTeamDriverContractByIDQuery,
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
	teamDriverContract, teamDriverContractErr := teamDriverContractRepository.scanFromRows(rows)
	if teamDriverContractErr != nil {
		return nil, teamDriverContractErr
	}
	return teamDriverContract, nil
}

func (teamDriverContractRepository *TeamDriverContractRepository) FindCurrentByTeamID(id int) ([]*model.TeamDriverContract, error) {
	findCurrentByTeamIDQuery := "SELECT id FROM team_driver_contracts WHERE team_id = $1 AND CURRENT_DATE BETWEEN from_date AND to_date;"

	rows, rowsErr := teamDriverContractRepository.store.db.Query(
		findCurrentByTeamIDQuery,
		id,
	)
	if rowsErr != nil {
		if rowsErr == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}

		return nil, rowsErr
	}
	defer rows.Close()

	var contracts []*model.TeamDriverContract

	for rows.Next() {
		teamDriverContract, err := teamDriverContractRepository.scanFromRows(rows)
		if err != nil {
			return contracts, err
		}

		contracts = append(contracts, teamDriverContract)
	}

	if err := rows.Err(); err != nil {
		return contracts, err
	}

	return contracts, nil
}

func (teamDriverContractRepository *TeamDriverContractRepository) FindCurrentByDriverID(id int) (*model.TeamDriverContract, error) {
	findDriverByIDQuery := "SELECT id FROM team_driver_contracts WHERE driver_id = $1 AND CURRENT_DATE BETWEEN from_date AND to_date;"
	var contractID int
	if err := teamDriverContractRepository.store.db.QueryRow(
		findDriverByIDQuery,
		id,
	).Scan(&contractID); err != nil {
		if err == sql.ErrNoRows {
			return nil, store.ErrRecordNotFound
		}
		return nil, err
	}

	return teamDriverContractRepository.Find(contractID)
}

func (teamDriverContractRepository *TeamDriverContractRepository) Update(teamDriverContract *model.TeamDriverContract) error {
	updateTeamDriverContractQuery := "UPDATE team_driver_contracts " +
		"SET from_date = $2, " +
		"to_date = $3, " +
		"team_id = $4 " +
		"driver_id = $5, " +
		"WHERE id = $1;"

	countResult, countResultErr := teamDriverContractRepository.store.db.Exec(
		updateTeamDriverContractQuery,
		teamDriverContract.ID,
		teamDriverContract.FromDate,
		teamDriverContract.ToDate,
		teamDriverContract.Team.ID,
		teamDriverContract.Driver.ID,
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

func (teamDriverContractRepository *TeamDriverContractRepository) Delete(id int) error {
	deleteTeamDriverContractQuery := "DELETE FROM team_driver_contracts WHERE id = $1;"

	countResult, countResultErr := teamDriverContractRepository.store.db.Exec(
		deleteTeamDriverContractQuery,
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

func (teamDriverContractRepository *TeamDriverContractRepository) scanFromRows(rows *sql.Rows) (*model.TeamDriverContract, error) {
	var (
		teamDriverContract *model.TeamDriverContract
		teamID             int
		driverID           int
	)

	if err := rows.Scan(
		&teamDriverContract.ID,
		&teamDriverContract.FromDate,
		&teamDriverContract.ToDate,
		&teamID,
		&driverID,
	); err != nil {
		return nil, err
	}

	tmpTeam, teamErr := teamDriverContractRepository.store.Team().Find(teamID)
	if teamErr != nil {
		return nil, teamErr
	}

	tmpDriver, driverErr := teamDriverContractRepository.store.Driver().Find(driverID)
	if driverErr != nil {
		return nil, driverErr
	}

	teamDriverContract.Team = tmpTeam
	teamDriverContract.Driver = tmpDriver

	return teamDriverContract, nil
}
