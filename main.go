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

	// ! Want to Upload files? You can use UploadThing on your Routes
	// * Validate if your .env file has the correct configuration
	// upload, err := uploader.UseUploader()
	// * Create a PrepareUpload object with customs values
	// upload.PrepareUpload()
	// * Upload the file you Prepared
	// upload.UploadFile()

	authHandler := auth.NewHandler()
	authHandler.BuildRoutes(router)

	userHandler := user.NewHandler()
	userHandler.BuildRoutes(router)

	logs.OkLog(fmt.Sprintf("Listening on port %s", s.addr))

	return http.ListenAndServe(s.addr, router.Router)
}
