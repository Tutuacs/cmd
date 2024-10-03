package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/routes"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewApiServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := routes.NewRouter()

	/*
		TODO:
		Set the handlers and pass the router to build the routes
	*/

	logs.OkLog(fmt.Sprintf("Listening on port %s", s.addr))

	return http.ListenAndServe(s.addr, router.Router)
}
