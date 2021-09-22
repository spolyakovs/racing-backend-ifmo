package store

func (store *Store) createTables() error {
	if err := store.createTableUsers(); err != nil {
		return err
	}

	if err := store.createTableTeams(); err != nil {
		return err
	}

	if err := store.createTableDrivers(); err != nil {
		return err
	}

	if err := store.createTableRaces(); err != nil {
		return err
	}

	if err := store.createTableTeamDriverContracts(); err != nil {
		return err
	}

	if err := store.createTableRaceResults(); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableUsers() error {
	createTableUsersQuery := "CREATE TABLE IF NOT EXISTS users (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"username varchar NOT NULL UNIQUE," +
		"email varchar NOT NULL UNIQUE," +
		"encrypted_password varchar NOT NULL );"

	if _, err := store.db.Exec(createTableUsersQuery); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableTeams() error {
	createTableTeamsQuery := "CREATE TABLE IF NOT EXISTS teams (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"name varchar NOT NULL UNIQUE," +
		"engine_manufacturer varchar NOT NULL );"

	if _, err := store.db.Exec(createTableTeamsQuery); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableDrivers() error {
	createTableDriversQuery := "CREATE TABLE IF NOT EXISTS drivers (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"first_name varchar NOT NULL," +
		"last_name varchar NOT NULL," +
		"birth_date date NOT NULL );"

	if _, err := store.db.Exec(createTableDriversQuery); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableRaces() error {
	createTableRacesQuery := "CREATE TABLE IF NOT EXISTS races (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"name varchar NOT NULL," +
		"location varchar NOT NULL," +
		"date date NOT NULL );"

	if _, err := store.db.Exec(createTableRacesQuery); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableTeamDriverContracts() error {
	createTableTeamDriverContractsQuery := "CREATE TABLE IF NOT EXISTS team_driver_contracts (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"from_date date NOT NULL," +
		"to_date date," +
		"team_id bigserial NOT NULL REFERENCES teams (id)," +
		"driver_id bigserial NOT NULL REFERENCES drivers (id) ON DELETE CASCADE );"

	if _, err := store.db.Exec(createTableTeamDriverContractsQuery); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableRaceResults() error {
	createTableRaceResultsQuery := "CREATE TABLE IF NOT EXISTS race_results (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"place integer NOT NULL UNIQUE," +
		"points integer NOT NULL UNIQUE DEFAULT 0," +
		"race_id bigserial NOT NULL REFERENCES races (id) ON DELETE CASCADE," +
		"driver_id bigserial NOT NULL REFERENCES drivers (id) ON DELETE CASCADE );"

	if _, err := store.db.Exec(createTableRaceResultsQuery); err != nil {
		return err
	}

	return nil
}
