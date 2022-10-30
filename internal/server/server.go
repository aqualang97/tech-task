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
	err = s.createTables()
	if err != nil {
		log.Println(err)
	}
	err = http.ListenAndServe(s.cfg.PortListen, r)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func (s *Server) createTables() error {
	_, err := s.db.Exec("create table if not exists data(" +
		"transaction_id      integer          not null        unique," +
		"request_id          integer          not null," +
		"terminal_id         integer          not null," +
		"partner_object_id   integer          not null," +
		"amount_total        integer          not null," +
		"amount_original     integer          not null," +
		"commission_ps       double precision not null," +
		"commission_client   integer          not null," +
		"commission_provider double precision not null," +
		"date_input          timestamp        not null," +
		"date_post           timestamp        not null," +
		"status              varchar(10)      not null," +
		"payment_type        varchar(10)      not null," +
		"payment_number      varchar(15)      not null," +
		"service_id          integer          not null," +
		"service             varchar(20)      not null," +
		"payee_id            integer          not null," +
		"payee_name          varchar(10)      not null," +
		"payee_bank_mfo      integer          not null," +
		"payee_bank_account  text             not null," +
		"payment_narrative   text             not null" +
		")")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = s.db.Exec("alter table data owner to " + s.cfg.PGUserName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
