package server

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"techTask/config"
	"techTask/internal/controller"
	"techTask/internal/models"
	"techTask/internal/router"
	"techTask/internal/services"
	"techTask/internal/store"
	"techTask/internal/worker_pool"
	"time"
)

type Server struct {
	cfg *config.Config
	db  *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Start(filePath string) error {
	db := s.db
	TX, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	storage := store.NewStore(db)

	service, err := services.NewManager(storage)

	controller := controller.NewController(service, s.cfg)
	mux := http.NewServeMux()

	router.Router(controller, mux)

	err = s.StartParse(filePath, db, TX)
	if err != nil {
		log.Println(err)
	}

	err = http.ListenAndServe(s.cfg.PortListen, mux)

	return nil
}

func (s *Server) StartParse(filePath string, db *sql.DB, TX *sql.Tx) error {
	pool := worker_pool.NewPool(4)
	count := pool.Count
	wg := &sync.WaitGroup{}

	file, err := os.Open(filePath)
	if err != nil {
		log.Println("can't open file; err:", err)
		return err
	}
	readFile, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println("can't read file; err:", err)
		return err
	}
	for i := 0; i < count; i++ {
		go func(i int) {
			pool.Start(filePath, wg, i, db, TX)
		}(i)
		wg.Add(1)

	}

	for i := 1; i < len(readFile); i++ {
		var data models.DataParse

		data.TransactionID, _ = strconv.Atoi(readFile[i][0])
		data.RequestID, _ = strconv.Atoi(readFile[i][1])
		data.TerminalID, _ = strconv.Atoi(readFile[i][2])
		data.PartnerObjectID, _ = strconv.Atoi(readFile[i][3])
		data.AmountTotal, _ = strconv.Atoi(readFile[i][4])
		data.AmountOriginal, _ = strconv.Atoi(readFile[i][5])
		data.CommissionPS, _ = strconv.ParseFloat(readFile[i][6], 64)
		data.CommissionClient, _ = strconv.Atoi(readFile[i][7])
		data.CommissionProvider, _ = strconv.ParseFloat(readFile[i][8], 64)
		data.DateInput, _ = time.Parse("2006-01-02 15:04:05", readFile[i][9])
		data.DatePost, _ = time.Parse("2006-01-02 15:04:05", readFile[i][10])
		data.Status = readFile[i][11]
		data.PaymentType = readFile[i][12]
		data.PaymentNumber = readFile[i][13]
		data.ServiceID, _ = strconv.Atoi(readFile[i][14])
		data.Service = readFile[i][15]
		data.PayeeID, _ = strconv.Atoi(readFile[i][16])
		data.PayeeName = readFile[i][17]
		data.PayeeBankMfo, _ = strconv.Atoi(readFile[i][18])
		data.PayeeBankAccount = readFile[i][19]
		data.PaymentNarrative = readFile[i][20]

		pool.StartSendData <- &data
	}

	pool.Stop()

	wg.Wait()
	return nil
}
