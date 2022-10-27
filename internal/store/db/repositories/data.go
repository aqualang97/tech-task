package repositories

import "database/sql"

type DataRepo struct {
	DB *sql.DB
	TX *sql.Tx
}

func NewDataRepo(db *sql.DB) *DataRepo {
	return &DataRepo{
		DB: db,
	}
}
