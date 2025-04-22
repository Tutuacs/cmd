package api

import (
	"fmt"
	"net/http"

	"github.com/Tutuacs/internal/auth"
	"github.com/Tutuacs/internal/user"
	"github.com/Tutuacs/pkg/db"
	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/routes"
	"github.com/rs/cors" // Import the CORS package
)

type APIServer struct {
	addr int64
}

func NewApiServer(addr int64) (*APIServer, error) {
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

	// * Create a Db Connection to use on that full application
	conn, err := db.NewConnection()
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error connecting to the database: %s", err))
		return err
	}

	authStore, _ := auth.NewStore(conn)
	authHandler := auth.NewHandler(authStore)
	authHandler.BuildRoutes(router)

	userStore, _ := user.NewStore(conn)
	userHandler := user.NewHandler(userStore)
	userHandler.BuildRoutes(router)

	// Enable CORS
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Allow all origins (you can specify specific origins instead)
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,  // Allow credentials (e.g., cookies)
		Debug:            false, // Enable debug logging for CORS (optional)
	})

	// Wrap the router with the CORS middleware
	handler := corsMiddleware.Handler(router.Router)

	logs.OkLog(fmt.Sprintf("Listening on port :%d", s.addr))

	// Start the server with the CORS-enabled handler
	return http.ListenAndServe(fmt.Sprintf(":%d", s.addr), handler)
}
