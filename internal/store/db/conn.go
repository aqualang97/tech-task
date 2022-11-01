package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"techTask/config"
	"time"
)

func Conn() (*sql.DB, error) {
	cfg := config.NewConfig()

	host := cfg.PGHOST
	port := cfg.PGPort
	user := cfg.PGUserName
	password := cfg.PGPassword
	dbname := cfg.PGDBName
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println("psqlInfo", psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(2)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(10 * time.Second)
	return db, nil
}
