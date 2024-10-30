package api

import (
	"fmt"
	"net/http"

	"github.com/Tutuacs/internal/auth"
	"github.com/Tutuacs/internal/user"
	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/routes"
)

type APIServer struct {
	addr string
}

func NewApiServer(addr string) (*APIServer, error) {
	return &APIServer{
		addr: addr,
	}, nil
}

func (s *APIServer) Run() error {
	router := routes.NewRouter()

	// ! Want to use WebSocket? Uncomment the following lines
	// * Create hanldersFunctions on the pkg/ws package
	// wsHandler := ws.NewWsHandler()
	// wsHandler.BuildRoutes(router)

	authHandler := auth.NewHandler()
	authHandler.BuildRoutes(router)

	userHandler := user.NewHandler()
	userHandler.BuildRoutes(router)

	logs.OkLog(fmt.Sprintf("Listening on port %s", s.addr))

	return http.ListenAndServe(s.addr, router.Router)
}
