package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/spolyakovs/racing-backend-ifmo/internal/app/store"
)

func Start(config *Config) error {
	db, dbErr := newDB(config.DatabaseURL)
	if dbErr != nil {
		return dbErr
	}

	defer db.Close()

	store, storeErr := store.New(db)
	if storeErr != nil {
		return storeErr
	}

	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	return http.ListenAndServe(config.BindAddr, srv)
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
