package server

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"techTask/config"
	"techTask/internal/controller"
	"techTask/internal/router"
	"techTask/internal/services"
	"techTask/internal/store"
)

type Server struct {
	cfg *config.Config
	db  *sql.DB
}

func NewServer(cfg *config.Config, db *sql.DB) *Server {
	return &Server{cfg: cfg, db: db}
}

func (s *Server) Start() error {
	db := s.db
	//TX, err := db.Begin()
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	storage := store.NewStore(s.cfg, db)

	service, err := services.NewManager(storage)

	ctr := controller.NewController(service, s.cfg)
	r := gin.Default()
	router.Router(ctr, r)
	if err != nil {
		log.Println(err)
	}

	err = http.ListenAndServe(s.cfg.PortListen, r)

	return nil
}
