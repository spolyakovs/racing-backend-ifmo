package store

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	config         *Config
	db             *sql.DB
	userRepository *UserRepository
}

func New(config *Config) *Store {
	return &Store{
		config: config,
	}
}

// TODO: move creating and dropping tables to separate functions and
// execute them when starting and shutting down server or tests

func (store *Store) Open() error {
	db, err := sql.Open("postgres", store.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	store.db = db

	if err := store.createTables(); err != nil {
		return err
	}

	return nil
}

func (store *Store) Close() {
	store.db.Close()
}

func (store *Store) User() *UserRepository {
	if store.userRepository != nil {
		return store.userRepository
	}

	store.userRepository = &UserRepository{
		store: store,
	}

	return store.userRepository
}

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
	create_table_users_query := "CREATE TABLE IF NOT EXISTS users (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"username varchar NOT NULL UNIQUE," +
		"email varchar NOT NULL UNIQUE," +
		"encrypted_password varchar NOT NULL );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableTeams() error {
	create_table_users_query := "CREATE TABLE IF NOT EXISTS teams (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"name varchar NOT NULL UNIQUE," +
		"engine_manufacturer varchar NOT NULL );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableDrivers() error {
	create_table_users_query := "CREATE TABLE IF NOT EXISTS drivers (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"first_name varchar NOT NULL," +
		"last_name varchar NOT NULL," +
		"birth_date date NOT NULL );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableRaces() error {
	create_table_users_query := "CREATE TABLE IF NOT EXISTS races (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"name varchar NOT NULL," +
		"location varchar NOT NULL," +
		"date date NOT NULL );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableTeamDriverContracts() error {
	create_table_users_query := "CREATE TABLE IF NOT EXISTS team_driver_contracts (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"from_date date NOT NULL," +
		"to_date date," +
		"team_id bigserial NOT NULL REFERENCES teams (id)," +
		"driver_id bigserial NOT NULL REFERENCES drivers (id) ON DELETE CASCADE );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}

func (store *Store) createTableRaceResults() error {
	create_table_users_query := "CREATE TABLE IF NOT EXISTS race_results (" +
		"id bigserial NOT NULL PRIMARY KEY," +
		"place integer NOT NULL UNIQUE," +
		"points integer NOT NULL UNIQUE DEFAULT 0," +
		"race_id bigserial NOT NULL REFERENCES races (id) ON DELETE CASCADE," +
		"driver_id bigserial NOT NULL REFERENCES drivers (id) ON DELETE CASCADE );"

	if _, err := store.db.Exec(create_table_users_query); err != nil {
		return err
	}

	return nil
}
