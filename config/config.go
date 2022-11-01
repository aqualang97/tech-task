package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	PGUserName string
	PGPassword string
	PGDBName   string
	PGHOST     string
	PGPort     string
	SSLMode    string
	//PGMaxIdleCons     int
	//PGMaxOpenCons     int
	//PGConsMaxLifeTime int
	PortListen string
}

func NewConfig() *Config {
	cfg := &Config{}
	err := godotenv.Load()
	if err != nil {
		log.Println("Can`t load .env")
	}
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println("Can`t init configs")
	}

	cfg.PGPassword = os.Getenv("DB_PASSWORD")
	cfg.PGUserName = viper.GetString("db.username")
	cfg.PGDBName = viper.GetString("db.dbname")
	cfg.PGHOST = viper.GetString("db.host")
	cfg.PGPort = viper.GetString("db.port")
	cfg.SSLMode = viper.GetString("db.sslmode")
	cfg.PortListen = viper.GetString("port")

	//err = godotenv.Load("config/config.env")
	//if err != nil {
	//	log.Fatal("Can`t load .env")
	//}
	//cfg.PGUserName = os.Getenv("USER")
	//cfg.PGDBName = os.Getenv("DBNAME")
	//
	//cfg.PGHOST = os.Getenv("host")
	//cfg.PGPort, _ = strconv.Atoi(os.Getenv("port"))
	//cfg.PGMaxIdleCons, _ = strconv.Atoi(os.Getenv("maxIdleConns"))
	//cfg.PGMaxOpenCons, _ = strconv.Atoi(os.Getenv("maxOpenConns"))
	//cfg.PGConsMaxLifeTime, _ = strconv.Atoi(os.Getenv("connMaxLifetime"))
	//cfg.PortListen = os.Getenv("portListen")
	fmt.Println(cfg)

	return cfg
}
