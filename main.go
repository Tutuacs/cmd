package main

import (
	"fmt"

	"github.com/Tutuacs/pkg/config"
	"github.com/Tutuacs/pkg/logs"

	"github.com/Tutuacs/cmd/api"
)

func main() {
	conf_API := config.GetAPI()

	server, err := api.NewApiServer(conf_API.Port, "127.0.0.1:6379")
	if err != nil {
		logs.ErrorLog(fmt.Sprintf("Error creating server: %s", err))
		return
	}

	if err := server.Run(); err != nil {
		logs.ErrorLog(fmt.Sprintf("Error starting server: %s", err))
	}
}
