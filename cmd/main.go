package main

import (
	"log"
	"os"
	"techTask/config"

	"techTask/internal/server"
	db2 "techTask/internal/store/db"
)

func main() {

	db, err := db2.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	cfg := config.NewConfig()
	s := server.NewServer(cfg, db)
	err = s.Start()

	if err != nil {
		log.Println("Can't start server with error", err)
		os.Exit(1)
	}

}
