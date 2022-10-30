package store

import (
	"database/sql"
	"techTask/config"
	"techTask/internal/store/db/repositories"
)

type Store struct {
	Data *repositories.DataRepo
}

func NewStore(cfg *config.Config, db *sql.DB) *Store {
	dr := repositories.NewDataRepo(db)
	return &Store{Data: dr}
}
