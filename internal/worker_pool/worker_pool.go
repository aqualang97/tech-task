package worker_pool

import (
	"database/sql"
	"sync"
	"techTask/internal/models"
	"techTask/internal/parser"
)

type WorkerPool struct {
	Count         int
	StartSendData chan *models.DataParse
	StopSend      chan bool
}

func NewPool(count int) *WorkerPool {
	return &WorkerPool{
		Count:         count,
		StartSendData: make(chan *models.DataParse),
		StopSend:      make(chan bool),
	}
}

func (pool *WorkerPool) Stop() {
	for i := 0; i < pool.Count; i++ {
		pool.StopSend <- false
	}
}

func (pool *WorkerPool) Start(filePath string, wg *sync.WaitGroup, goNum int, db *sql.DB, TX *sql.Tx) {
	var data *models.DataParse

	defer wg.Done()
	for {
		select {
		case data = <-pool.StartSendData:
			_ = parser.ParseFromFile(data, goNum, db, TX)
		case <-pool.StopSend:
			return
		}
	}
}
