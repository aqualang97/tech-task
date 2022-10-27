package store

import (
	"database/sql"
	"techTask/internal/store/db/repositories"
)

type Store struct {
	Data *repositories.DataRepo
}

func NewStore(db *sql.DB) *Store {
	dr := repositories.NewDataRepo(db)
	return &Store{Data: dr}
}
