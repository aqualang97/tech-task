package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

type Config struct {
	PGUserName        string
	PGPassword        string
	PGDBName          string
	PGHOST            string
	PGPort            int
	PGMaxIdleCons     int
	PGMaxOpenCons     int
	PGConsMaxLifeTime int
	PortListen        string
}

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load("config/postgres.env")
	if err != nil {
		log.Fatal("Can`t load postgres.env")
	}
	cfg.PGPassword = os.Getenv("PG_PASSWORD")

	err = godotenv.Load("config/config.env")
	if err != nil {
		log.Fatal("Can`t load postgres.env")
	}
	cfg.PGUserName = os.Getenv("USER")
	cfg.PGDBName = os.Getenv("DBNAME")

	cfg.PGHOST = os.Getenv("host")
	cfg.PGPort, _ = strconv.Atoi(os.Getenv("port"))
	cfg.PGMaxIdleCons, _ = strconv.Atoi(os.Getenv("maxIdleConns"))
	cfg.PGMaxOpenCons, _ = strconv.Atoi(os.Getenv("maxOpenConns"))
	cfg.PGConsMaxLifeTime, _ = strconv.Atoi(os.Getenv("connMaxLifetime"))
	cfg.PortListen = os.Getenv("portListen")
	fmt.Println(cfg)

	return cfg
}
