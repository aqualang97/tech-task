package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"techTask/config"
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
	//db.SetMaxIdleConns(cfg.PGMaxIdleCons)
	//db.SetMaxOpenConns(cfg.PGMaxOpenCons)
	//db.SetConnMaxLifetime(time.Duration(cfg.PGConsMaxLifeTime) * time.Second)
	return db, nil
}
