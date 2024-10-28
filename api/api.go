package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Tutuacs/internal/auth"
	"github.com/Tutuacs/internal/user"
	"github.com/Tutuacs/pkg/cache"
	"github.com/Tutuacs/pkg/logs"
	"github.com/Tutuacs/pkg/routes"
	"github.com/Tutuacs/pkg/types"
)

type APIServer struct {
	addr   string
	pubsub *cache.RedisPubSub
}

func NewApiServer(addr, redisAddr string) (*APIServer, error) {
	// Inicializa o RedisPubSub
	pubsub, err := cache.NewRedisPubSub(redisAddr)
	if err != nil {
		return nil, fmt.Errorf("erro ao inicializar o RedisPubSub: %w", err)
	}

	// Adiciona os manipuladores para os tópicos de exemplo
	pubsub.Subscribe("/hello", cache.HandleHello)
	pubsub.Subscribe("/unknown", cache.HandleUnknown)

	return &APIServer{
		addr:   addr,
		pubsub: pubsub,
	}, nil
}

func (s *APIServer) Run() error {
	router := routes.NewRouter()

	// Configura os manipuladores de rotas existentes
	authHandler := auth.NewHandler()
	authHandler.BuildRoutes(router)

	userHandler := user.NewHandler()
	userHandler.BuildRoutes(router)

	// Adiciona uma rota para publicar mensagens de exemplo
	router.Router.HandleFunc("/publish", s.handlePublish)

	logs.OkLog(fmt.Sprintf("Listening on port %s", s.addr))

	// Inicia o listener do PubSub em uma goroutine
	go s.pubsub.Listen()

	return http.ListenAndServe(s.addr, router.Router)
}

// Manipulador de exemplo para publicar uma mensagem
func (s *APIServer) handlePublish(w http.ResponseWriter, r *http.Request) {
	var msg types.Message
	// Decodifica a mensagem do corpo da solicitação
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid message format", http.StatusBadRequest)
		return
	}

	// Publica a mensagem no tópico especificado
	if err := s.pubsub.Publish(context.Background(), msg.Topic, msg); err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Message published to topic: %s", msg.Topic)
}
