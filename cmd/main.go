package main

import (
	"log"
	"os"
	"path/filepath"
	"techTask/config"
	"techTask/internal/parser"
	"techTask/internal/server"
	db2 "techTask/internal/store/db"
)

func main() {

	url := "https://drive.google.com/u/0/uc?id=1IwZ3uUCHGpSL2OoQu4mtbw7Ew3ZamcGB&export=download"
	filename := "dataParse.csv"

	db, err := db2.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	p := parser.NewDataParser(db)

	err = p.DownloadFile(url, filename)
	path, _ := filepath.Abs(filename)

	cfg := config.NewConfig()
	s := server.NewServer(cfg, db)
	err = s.Start(path)

	if err != nil {
		log.Println("Can't start server with error", err)
		os.Exit(1)
	}

	p.SaveToDB(filename)

}
