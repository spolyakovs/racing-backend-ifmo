package store

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/spolyakovs/racing-backend-ifmo/internal/app/model"
)

type Repository interface {
	handleModel()
	CreateDefault(model.Model, string, ...string) error
	FindByFieldDefault(string, interface{}, string, ...string) (model.Model, error)
	UpdateDefault(model.Model, string, ...string) error
	Delete(string, int) error
}

type DefaultRepository struct {
	store *Store
}

func (defaultRepository *DefaultRepository) CreateDefault(mod model.Model, tableName string, queryArgs ...string) error {
	queryValues := []interface{}{}

	switch m := mod.(type) {
	case *model.User:

		query := fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES (%s) RETURNING id",
			tableName,
			makeArgsString(queryArgs...),
			makeValsString(queryArgs...),
		)

		reflectValue := reflect.ValueOf(m).Elem()

		for i := 0; i < reflectValue.NumField() && i < len(queryArgs); i++ {
			field := reflectValue.Field(i)
			if field.IsValid() && field.CanInterface() {
				queryValues = append(queryValues, field.Interface())
			}
		}

		if err := defaultRepository.store.db.QueryRow(
			query,
			queryValues...,
		).Scan(&m.ID); err != nil {
			if err == sql.ErrNoRows {
				return ErrRecordNotFound
			}

			return err
		}

		return nil
	default:
		return ErrWrongModel
	}
}

func (defaultRepository *DefaultRepository) FindByFieldDefault(field string, value interface{}, tableName string, queryArgs ...string) (model.Model, error) {
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE %s = $1",
		makeArgsString(queryArgs...),
		tableName,
		field,
	)

	var mod model.Model

	switch tableName {
	case "users":
		mod = &model.User{}
	// case "teams":
	// 	mod = &model.User{}
	// case "drivers":
	// 	mod = &model.User{}
	// case "races":
	// 	mod = &model.User{}
	// case "race_results":
	// 	mod = &model.User{}
	// case "team_driver_contracts":
	// 	mod = &model.User{}
	default:
		return nil, ErrWrongTable
	}

	if err := model.ScanModelFromRow(
		defaultRepository.store.db.QueryRow(
			query,
			value,
		),
		mod.GetFields(),
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrRecordNotFound
		}

		return nil, err
	}

	return mod, nil
}

func (defaultRepository *DefaultRepository) UpdateDefault(m model.Model, tableName string, queryArgs ...string) error {
	switch mod := m.(type) {
	case *model.User:

		query := fmt.Sprintf(
			"UPDATE %s SET %s WHERE id = $1",
			tableName,
			makeUpdateSetString(queryArgs...),
		)

		reflectValue := reflect.ValueOf(mod).Elem()

		queryVals := []interface{}{mod.ID}

		for i := 0; i < reflectValue.NumField() && i < len(queryArgs); i++ {
			field := reflectValue.Field(i)
			if field.IsValid() && field.CanInterface() {
				queryVals = append(queryVals, field.Interface())
			}
		}

		countResult, countResultErr := defaultRepository.store.db.Exec(
			query,
			queryVals...,
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
	default:
		return ErrWrongModel
	}
}

func (defaultRepository *DefaultRepository) DeleteDefault(tableName string, id int) error {

	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableName)

	countResult, countResultErr := defaultRepository.store.db.Exec(
		query,
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
