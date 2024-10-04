package api

import (
	"fmt"
	"net/http"

	"github.com/Tutuacs/internal/eg"
	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/routes"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := routes.NewRouter()

	/*
		TODO:
		Set the handlers and pass the router to build the routes

	*/

	egHandler := eg.NewHandler()
	egHandler.BuildRoutes(router)

	logs.OkLog(fmt.Sprintf("Listening on port %s", s.addr))

	return http.ListenAndServe(s.addr, router.Router)
}
