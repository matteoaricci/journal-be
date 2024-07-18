package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/matteoaricci/journal-be/service/journal"
	"github.com/matteoaricci/journal-be/service/user"
)

type APIServer struct {
	addr string 
	db *sql.DB 
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subRouter)

	journalHandler := journal.NewHandler()
	journalHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}