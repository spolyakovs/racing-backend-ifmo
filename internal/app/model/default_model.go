package model

import (
	"database/sql"
)

type Model interface {
	GetFields() []interface{}
}

type DefaultModel struct {
	ID int `json:"id"`
}

func (defaultModel *DefaultModel) GetFields() []interface{} {
	return []interface{}{&defaultModel.ID}
}

func ScanModelFromRow(row *sql.Row, fields []interface{}) error {
	if err := row.Scan(fields...); err != nil {
		if err == sql.ErrNoRows {
			return ErrRecordNotFound
		}

		return err
	}

	return nil
}
