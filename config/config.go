package config

import (
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

	//fmt.Println(cfg)

	return cfg
}
