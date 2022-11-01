package main

import (
	"log"
	"os"
	"techTask/config"

	"techTask/internal/server"
	db "techTask/internal/store/db"
)

func main() {

	d, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer d.Close()
	cfg := config.NewConfig()
	s := server.NewServer(cfg, d)
	err = s.Start()

	if err != nil {
		log.Println("Can't start server with error", err)
		os.Exit(1)
	}

}
